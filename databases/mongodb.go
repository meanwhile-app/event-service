package databases

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/meanwhile-app/event-service/configs"
)

var dbClient *mongo.Client

func ConnectMongoDB() *mongo.Client {
	if dbClient != nil {
		return dbClient
	}
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
		log.Fatal(err)
	}
	dbClient = client
	return client
}

func GetMongoDbClient() *mongo.Client {
	return dbClient
}
