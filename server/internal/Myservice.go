// server/internal/myservice.go

package internal

import (
	"context"
	"fmt"

	"github.com/Kadoshima/grpc-container/pb"
	// ↑ pbパッケージへのimportパスを正しく合わせてください。
	//   たとえば "github.com/Kadoshima/grpc-container/pb" 等。
	// 他に必要なら "google.golang.org/grpc/status" などをimport
)

// MyServiceImpl は pb.MyServiceServer インターフェースを実装するサーバ
type MyServiceImpl struct {
	pb.UnimplementedMyServiceServer
	// ↑ これを埋め込むと「メソッド未実装」の警告を抑制できる
}

// DoSomething は .proto で定義したRPCメソッド
func (s *MyServiceImpl) DoSomething(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	// ここでサーバ側の処理を行う
	fmt.Printf("DoSomethingが呼ばれました: %s\n", req.UserId)

	// レスポンスを返す
	return &pb.Response{
		Result: "Hello, " + req.UserId,
	}, nil
}
