package app

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron"
	"grpc/internal/app/clientGRPC"
	"grpc/internal/repository"
	"grpc/internal/utils"
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
		_, errDB := db.CreateUser(context.Background(), *user)
		if errDB != nil {
			fmt.Println("Error creating user in DB: ", err)
			return
		}
		res, errGRPC := clientGRPC.CreateUser(context.Background(), user)
		if errGRPC != nil {
			fmt.Println("Error creating user with GRPC client: ", err)
			return
		}
		fmt.Println(res)
	})
	scheduler.StartAsync()
	scheduler.StartBlocking()
}
