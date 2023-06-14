package app

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc/internal/app/serverGRPC"
	"grpc/internal/repository"
	"grpc/internal/utils"
	"grpc/pb"
	"log"
	"net"
)

func Run() {
	//client, err := repository.NewMongo("user")

	//if err != nil {
	//	fmt.Errorf("repository.NewMongo", err)
	//}

	//shed := gocron.NewScheduler(time.UTC)
	//shed.Every(5).Seconds().Do(func() {
	//	//user := utils.GenerateUser()
	//})
	//db := client.Client.Collection("users")
	//db.InsertOne(context.Background(), "user")
	//shed.StartAsync()
	user := utils.GenerateUser()
	client, err := repository.Init(context.Background(), "user")
	db := repository.NewRepository(client.Collection("users"))

	db.Create(context.Background(), *user)

	gRPC := grpc.NewServer()

	pb.RegisterUserServiceServer(gRPC, &serverGRPC.Server{})
	reflection.Register(gRPC)

	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = gRPC.Serve(listener)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}
}
