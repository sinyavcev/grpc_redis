package app

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"reciever/internal/app/serverGRPC"
	"reciever/internal/repository"
	"reciever/pb"
)

func Run() {

	//client, err := repository.Init()
	//
	//db := repository.NewRepository(client.Collection("users"))
	//db.CreateUser(context.Background(), *user)

	gRPC := grpc.NewServer()
	client, _ := repository.Init(context.Background(), "localhost:6379")

	repository := repository.NewRepository(client)
	server := serverGRPC.Server{Repos: *repository}

	pb.RegisterUserServiceServer(gRPC, &server)
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
