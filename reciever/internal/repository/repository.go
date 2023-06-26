package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sinyavcev/proto/pb"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user pb.CreateUserRequest) error
	UpdateUser(ctx context.Context, user pb.UpdateUserRequest) (string, error)
}

type Repository struct {
	UserRepository UserRepository
}

func NewRepository(redis *redis.Client) *Repository {
	return &Repository{UserRepository: NewRedisMethod(redis)}
}
