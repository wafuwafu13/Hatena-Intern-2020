package grpc

import (
	"context"

	pb "github.com/wafuwafu13/Hatena-Intern-2020/services/fetcher-go/pb/fetcher"
	"github.com/wafuwafu13/Hatena-Intern-2020/services/fetcher-go/fetcher"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Server は pb.RendererServer に対する実装
type Server struct {
	pb.UnimplementedFetcherServer
	healthpb.UnimplementedHealthServer
}

// NewServer は gRPC サーバーを作成する
func NewServer() *Server {
	return &Server{}
}

// Fetch は受け取ったURLからタイトルを取得して返す
func (s *Server) Fetch(ctx context.Context, in *pb.FetcherRequest) (*pb.FetcherReply, error) {
	title, err := fetcher.Fetch(ctx, in.Src)
	if err != nil {
		return nil, err
	}
	return &pb.FetcherReply{Title: title}, nil
}
