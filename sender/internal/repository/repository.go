package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"grpc/pb"
)

type UserService interface {
	CreateUser(ctx context.Context, user pb.CreateUserRequest) (string, error)
}

type Repository struct {
	UserService
}

func NewRepository(db *mongo.Collection) *Repository {
	return &Repository{
		UserService: NewMongoMethod(db),
	}
}
