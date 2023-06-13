package app

import (
	"fmt"
	"github.com/go-co-op/gocron"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc/internal/app/serverGRPC"
	"grpc/internal/repository"
	"grpc/internal/service"
	"grpc/pb"
	"log"
	"net"
	"time"
)

func Run() {
	db, err := repository.NewMongo("user")
	if err != nil {
		fmt.Errorf("repository.NewMongo", err)
	}

	sed := gocron.NewScheduler(time.UTC)

	// Every starts the job immediately and then runs at the
	// specified interval
	sed.Every(5).Seconds().Do(func() {
		fmt.Println("Раз/2/123/213")
	})

	sed.StartAsync()

	repository := repository.NewRepository(db)
	usecases := service.NewUsecases(repository)

	s := grpc.NewServer()
	server := serverGRPC.NewServer(usecases)

	pb.RegisterArticleServiceServer(s, server)
	reflection.Register(s)

	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = s.Serve(listener)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}
}
