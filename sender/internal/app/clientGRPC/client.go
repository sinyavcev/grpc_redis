package clientGRPC

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/pb"
	"log"
	"time"
)

const (
	address = "localhost:8080"
)

func NewClient() pb.UserServiceClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	client := pb.NewUserServiceClient(conn)

	_, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*5000))
	defer cancel()

	return client
}
