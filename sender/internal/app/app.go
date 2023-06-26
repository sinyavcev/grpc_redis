package app

import (
	"context"
	"fmt"
	"github.com/go-co-op/gocron"
	"github.com/spf13/viper"
	"log"
	"os"
	"sender/configs"
	"sender/internal/app/clientGRPC"
	"sender/internal/common/logger"
	"sender/internal/common/mongo"
	"sender/internal/repository"
	"sender/internal/utils"
	"time"
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

	clientMongo, err := mongo.Init(
		context.Background(),
		viper.GetString("DB_ADDRESS"),
		viper.GetString("DB_NAME"),
	)
	if err != nil {
		fmt.Errorf("repository.Init %w", err)
	}

	var (
		clientGRPC = clientGRPC.NewClient(viper.GetString("CLIENT_ADDRESS"))
		coll       = clientMongo.Collection(viper.GetString("DB_COLLECTION"))
		db         = repository.NewRepository(coll)
		scheduler  = gocron.NewScheduler(time.UTC)
	)

	scheduler.Every(5).Seconds().Do(func() {
		user := utils.GenerateUser()
		_, errDB := db.UserRepository.CreateUser(context.Background(), *user)
		if errDB != nil {
			fmt.Errorf("error creating user in DB: %w", err)
			return
		}
		res, errGRPC := clientGRPC.CreateUser(context.Background(), user)
		if errGRPC != nil {
			fmt.Errorf("error creating user with GRPC client: %w", err)
			return
		}
		logger.Info(res)
	})
	scheduler.StartAsync()
	scheduler.StartBlocking()

}
