// server/cmd/main.go
package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Kadoshima/grpc-container/pb"
	"github.com/Kadoshima/grpc-container/server/internal"

	// ↑ internal/myservice.go を置いた場所に合わせる

	"google.golang.org/grpc"
)

func main() {
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v", port, err)
	}
	log.Printf("gRPC server listening on %s\n", port)

	grpcServer := grpc.NewServer()

	// ★ 追加： サービス実装のインスタンス生成
	myService := &internal.MyServiceImpl{}

	// ★ 追加： サーバに登録する
	pb.RegisterMyServiceServer(grpcServer, myService)

	// サーバ起動
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("Server started successfully")
}
