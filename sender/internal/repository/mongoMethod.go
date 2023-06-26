package repository

import (
	"context"
	"fmt"
	"github.com/sinyavcev/proto/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	collection *mongo.Collection
}

func (d *DB) CreateUser(ctx context.Context, user pb.CreateUserRequest) (string, error) {
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user due to error: %v", err)
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to convert object id to hex")
	}
	return oid.Hex(), nil
}

func NewMongoMethod(collection *mongo.Collection) *DB {
	return &DB{collection: collection}
}
