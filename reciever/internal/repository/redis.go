package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func Init(ctx context.Context, addr string, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set by default
		DB:       db,       // use default DB
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("Failed to connect Redis: %v", err)
	}

	return client, nil
}
