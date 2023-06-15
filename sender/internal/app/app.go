package app

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron"
	"grpc/internal/app/clientGRPC"
	"grpc/internal/repository"
	"grpc/internal/utils"
	"log"
	"time"
)

func Run() {
	clientGRPC := clientGRPC.NewClient()

	clientMongo, err := repository.Init(context.Background(), "user")
	coll := clientMongo.Collection("users")
	db := repository.NewRepository(coll)

	if err != nil {
		fmt.Errorf("repository.Init %w", err)
	}

	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(5).Seconds().Do(func() {
		user := utils.GenerateUser()
		db.CreateUser(context.Background(), *user)
		res, err := clientGRPC.CreateUser(context.Background(), user)
		if err != nil {
			log.Fatalf("CreateUser: %v", err)
		}
		fmt.Println(res)
	})
	scheduler.StartAsync()
	scheduler.StartBlocking()
}
