package serverGRPC

import (
	"context"
	"grpc/pb"
)

type Usecases interface {
	CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.Article, error)
}

type Server struct {
	usecases Usecases
	pb.UnimplementedArticleServiceServer
}

func NewServer(usecases Usecases) *Server {
	return &Server{usecases: usecases}
}
