package serverGRPC

import (
	"context"
	"fmt"
	"grpc/pb"
)

func (s *Server) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.Article, error) {
	fmt.Println(req)
	//s.usecases.CreateArticle()
	return &pb.Article{
		Id:      "asd",
		Title:   "asd",
		Content: "asd",
	}, nil
}
