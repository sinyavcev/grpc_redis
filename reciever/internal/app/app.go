package app

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"reciever/internal/app/serverGRPC"
	"reciever/pb"
)

func Run() {

	//client, err := repository.Init()
	//
	//db := repository.NewRepository(client.Collection("users"))
	//db.CreateUser(context.Background(), *user)

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
