package service

import (
	"context"
	"grpc/pb"
)

type Repository interface {
	CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.Article, error)
}

type Usecases struct {
	repository Repository
}

func NewUsecases(repository Repository) *Usecases {
	return &Usecases{repository}
}
