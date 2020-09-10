package grpc

import (
	"context"

	my_app "github.com/wafuwafu/Hatena-Intern-2020/services/renderer-go/app"
	pb "github.com/wafuwafu/Hatena-Intern-2020/services/renderer-go/pb/renderer"
	"github.com/wafuwafu/Hatena-Intern-2020/services/renderer-go/renderer"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

// Server は pb.RendererServer に対する実装
type Server struct {
	pb.UnimplementedRendererServer
	healthpb.UnimplementedHealthServer
	app my_app.App
}

// NewServer は gRPC サーバーを作成する
func NewServer(app my_app.App) *Server {
	s := &Server{}
	s.app = app
	return s
}

// Render は受け取った文書を HTML に変換する
func (s *Server) Render(ctx context.Context, in *pb.RenderRequest) (*pb.RenderReply, error) {
	html, err := renderer.Render(ctx, in.Src, s.app)
	if err != nil {
		return nil, err
	}
	return &pb.RenderReply{Html: html}, nil
}
