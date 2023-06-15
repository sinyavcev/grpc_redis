package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"reciever/pb"
)

type DB struct {
	redis *redis.Client
}

func (db *DB) CreateUser(ctx context.Context, user pb.CreateUserRequest) error {
	err := db.redis.Set(ctx, user.Phone, user.Name, 0)
	if err != nil {
		fmt.Errorf("Redis insert failed", err)
		return err.Err()
	}
	return nil
}
func (db *DB) UpdateUser(ctx context.Context, user pb.CreateUserRequest) error {
	err := db.redis.Set(ctx, user.Phone, user.Name, 0)
	if err != nil {
		fmt.Errorf("Redis update failed", err)
		return err.Err()
	}
	return nil
}

func NewRedisMethod(redis *redis.Client) *DB {
	return &DB{redis: redis}
}
