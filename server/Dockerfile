# server/Dockerfile

###########################################################
# (1) 開発用ステージ (development)
###########################################################
FROM golang:1.22 AS development

WORKDIR /app

# 1) まず go.mod / go.sum だけコピー
COPY go.mod go.sum ./
RUN go mod download

# 2) 残りのソースを全部コピー
COPY . .

# 3) 依存チェック (ここで server/cmd/main.go などを読み込む)
RUN go mod tidy

# 4) 開発モードは直接 go run
#    server/cmd/main.go が実際のエントリポイントなら、パスに注意
CMD ["go", "run", "./server/cmd/main.go"]


###########################################################
# (2) 本番用ステージ (production)
###########################################################
FROM golang:1.22 AS builder

WORKDIR /app

# 1) 依存情報をコピー＆ダウンロード
COPY go.mod go.sum ./
RUN go mod download

# 2) ソース全体をコピー
COPY . .

# 3) バイナリビルド
#    server/cmd/main.go をビルド対象に指定
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./server/cmd/main.go

# 4) distroless イメージで最終成果物のみ
FROM gcr.io/distroless/base-debian11 AS production

COPY --from=builder /server /server
EXPOSE 50051
ENTRYPOINT ["/server"]
