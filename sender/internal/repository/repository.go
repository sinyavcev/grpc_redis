package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"grpc/internal/models"
)

type UserService interface {
	Create(ctx context.Context, user models.User) (string, error)
}

type Repository struct {
	UserService
}

func NewRepository(db *mongo.Collection) *Repository {
	return &Repository{
		UserService: NewMongoMethod(db),
	}
}
