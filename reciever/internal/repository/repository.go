package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"reciever/pb"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user pb.CreateUserRequest) error
}

type Repository struct {
	UserRepository UserRepository
}

func NewRepository(redis *redis.Client) *Repository {
	return &Repository{UserRepository: NewRedisMethod(redis)}
}
