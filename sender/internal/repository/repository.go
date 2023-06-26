package repository

import (
	"context"
	"github.com/sinyavcev/proto/pb"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user pb.CreateUserRequest) (string, error)
}

type Repository struct {
	UserRepository UserRepository
}

func NewRepository(db *mongo.Collection) *Repository {
	return &Repository{
		UserRepository: NewMongoMethod(db),
	}
}
