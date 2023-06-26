package clientGRPC

import (
	"context"

	"github.com/sinyavcev/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func NewClient(clientAddress string) pb.UserServiceClient {
	conn, err := grpc.Dial(clientAddress, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	client := pb.NewUserServiceClient(conn)

	_, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*5000))
	defer cancel()

	return client
}
