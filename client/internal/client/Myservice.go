package client

import (
	"context"
	"fmt"
	"time"

	pb "github.com/Kadoshima/grpc-container/pb"
	"google.golang.org/grpc"
)

type MyClient struct {
	conn   *grpc.ClientConn
	client pb.MyServiceClient
}

// NewMyClient は gRPC 接続を行い、MyClient を生成
func NewMyClient(target string) (*MyClient, error) {
	// TLSを使わない簡易例
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("gRPC接続失敗: %w", err)
	}

	return &MyClient{
		conn:   conn,
		client: pb.NewMyServiceClient(conn),
	}, nil
}

// Close は gRPC 接続を閉じる
func (m *MyClient) Close() {
	if m.conn != nil {
		m.conn.Close()
	}
}

// DoSomething はサーバへRPCを実行して結果を取得
func (m *MyClient) DoSomething(userID string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.Request{UserId: userID}
	resp, err := m.client.DoSomething(ctx, req)
	if err != nil {
		return "", fmt.Errorf("DoSomething呼び出し失敗: %w", err)
	}
	return resp.Result, nil
}
