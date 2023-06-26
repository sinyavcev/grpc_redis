package app

import (
	"context"
	"fmt"
	"github.com/sinyavcev/proto/pb"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"reciever/common/logger"
	"reciever/common/redis"
	"reciever/configs"
	"reciever/internal/app/serverGRPC"
	"reciever/internal/repository"
)

func Run() {

	if err := configs.LoadConfig("configs"); err != nil {
		log.Fatalf("LoadConfig", err.Error())
	}

	logger, err := logger.NewLogger(viper.GetString("LOG_LEVEL"))
	if err != nil {
		log.Printf("logger.NewLogger: %w", err)
		os.Exit(1)
	}

	client, err := redis.NewRedis(context.Background(), viper.GetString("DB_ADDRESS"))
	if err != nil {
		fmt.Errorf("redis.NewRedis: %w", err)
	}

	var (
		gRPC       = grpc.NewServer()
		repository = repository.NewRepository(client)
		server     = serverGRPC.Server{Repos: repository, Logger: logger}
	)

	pb.RegisterUserServiceServer(gRPC, &server)
	reflection.Register(gRPC)

	listener, err := net.Listen("tcp", viper.GetString("SERVER_ADDRESS"))
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = gRPC.Serve(listener)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}
}
