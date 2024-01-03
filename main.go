package main

import (
	"github.com/gin-gonic/gin"
	"github.com/meanwhile-app/event-service/configs"
	"github.com/meanwhile-app/event-service/databases"
	"github.com/meanwhile-app/event-service/server"
)

func main() {
	env := configs.LoadEnv()
	gin.SetMode(env["GIN_MODE"])
	databases.ConnectMongoDB()
	server.InitRoute()
	server.Run()
}
