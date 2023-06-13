package repository

import (
	"context"
	"grpc/pb"
)

func NewRepository(mongo *Mongo) *Mongo {
	return &Mongo{mongo.Client,
		mongo.DatabaseName}
}

func (m *Mongo) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.Article, error) {
	return nil, nil
}
