# ===========================
# 1) ビルドステージ
# ===========================
FROM golang:1.22 AS builder

WORKDIR /app

# 依存情報を先にコピーしてキャッシュを活かす
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピー
COPY . .

# 必要であれば protoc をインストール（コメント解除して使う）
# RUN apt-get update && apt-get install -y protobuf-compiler

# Go向けのprotocプラグインをインストール（すでにdocker外で生成済なら不要）
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# .proto -> Goコード生成
RUN protoc \
  --proto_path=./proto \
  --go_out=pb --go_opt=paths=source_relative \
  --go-grpc_out=pb --go_grpc_opt=paths=source_relative \
  proto/myservice.proto

# バイナリをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o client-app ./cmd

# ===========================
# 2) 本番運用用ステージ (production)
# ===========================
FROM gcr.io/distroless/base-debian10 AS production

WORKDIR /app

# 上記 builder からバイナリだけコピー
COPY --from=builder /app/client-app /app/

# 起動コマンド（本番用）
CMD ["/app/client-app", "grpc-server:50052"]

# ===========================
# 3) 開発用ステージ (development)
# ===========================
FROM golang:1.22 AS development

WORKDIR /app

# 開発用はソースコードをコンテナ内にコピー or ホストからボリュームマウント
# ここでは最小限のファイルだけコピーし、後で docker-compose でマウント想定
COPY go.mod go.sum ./
RUN go mod download

# （必要なら）protoc を入れてもOK。ただし変更検知はホストマシン側かコンテナ再起動で実施
# RUN apt-get update && apt-get install -y protobuf-compiler

# 開発時の実行コマンド
#   - コンテナ起動時に go run し、ソースを直接実行
#   - ホスト側のファイル変更を反映させたい場合は、コンテナを再起動
CMD ["sh", "-c", "go run ./client/cmd/main.go grpc-server:50051"]
