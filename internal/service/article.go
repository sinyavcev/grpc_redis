package service

import (
	"context"
	"grpc/pb"
)

func (u *Usecases) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.Article, error) {
	return u.repository.CreateArticle(ctx, req)

}
