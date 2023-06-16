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
	err := db.redis.Set(ctx, user.Phone, user.Name, 0).Err()
	if err != nil {
		fmt.Errorf("Redis insert failed", err)
		return err
	}
	return nil
}
func (db *DB) UpdateUser(ctx context.Context, user pb.UpdateUserRequest) (string, error) {
	userName, err := db.redis.Get(ctx, user.Phone).Result()
	if err != nil {
		fmt.Errorf("db.redis.Get", err)
	}
	err = db.redis.Set(ctx, user.Phone, user.Name, 0).Err()
	if err != nil {
		fmt.Errorf("Redis update failed", err)
		return "", err
	}
	return userName, nil
}

func NewRedisMethod(redis *redis.Client) *DB {
	return &DB{redis: redis}
}
