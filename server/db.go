package server

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/nuntjw/go-gin-starter/configs"
)

func ConnectDB() {
	env := configs.GetEnv()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	credential := options.Credential{
		Username:   env["DB_USERNAME"],
		Password:   env["DB_PASSWORD"],
		AuthSource: env["DB_AUTH_SOURCE"],
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env["DB_URL"]).SetAuth(credential))
	if err != nil {
		println("3333")
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		println("4444")
		log.Fatal(err)
	}
}
