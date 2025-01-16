// client/cmd/main.go:
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Kadoshima/grpc-container/client/internal/client"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使い方: ./client-app <server-address:port>")
		os.Exit(1)
	}
	target := os.Args[1]

	c, err := client.NewMyClient(target)
	if err != nil {
		log.Fatalf("クライアント作成失敗: %v", err)
	}
	defer c.Close()

	result, err := c.DoSomething("test_user")
	if err != nil {
		log.Fatalf("RPC呼び出し失敗: %v", err)
	}
	log.Printf("サーバからの応答: %s", result)
}
