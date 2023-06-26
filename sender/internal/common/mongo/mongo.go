package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init(ctx context.Context, address string, database string) (db *mongo.Database, err error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(address))

	if err != nil {
		return nil, fmt.Errorf("failed to connect DB")
	}
	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("Failed ping")
	}

	return client.Database(database), nil
}
