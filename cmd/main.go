package main

import (
	"gateway/core/app"
	"gateway/core/context"
	"gateway/infrastructure/api/http"
	"gateway/infrastructure/storage/postgres"
	"gateway/infrastructure/storage/redis"
	"gateway/infrastructure/utils"
	"github.com/joho/godotenv"
	"log"
	//"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	utils.InitLogger("humo")
	redis.InitKvRedis()
}

func main() {

	//utils.Logger.Log("starting_service", "start...")

	err := postgres.InitConnection()

	if err != nil {
		//utils.Logger.Error("pg_init_connection", err)
		log.Fatal(err)
	}

	//aes := os.Getenv("AES")

	defer func() {
		err := postgres.CloseConnection()
		if err != nil {
			panic(err)
		}
	}()

	appService := app.ApplicationService{
		Context: context.ApplicationContext{
			//Logger: utils.Logger,
			//Kv:                      redis.KvDb,
			//HumoServiceProvider:     provider.HumoServiceProvider,
			//HumoSoapServiceProvider: provider.HumoSoapServiceProvider,
			//CryptoKey:               aes,
		},
	}

	server := http.CreateServer(&appService)
	err = server.Run(8000)
	if err != nil {
		log.Fatal(err)
		return
	}
}
