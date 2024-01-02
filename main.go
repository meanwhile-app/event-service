package main

import (
	"github.com/nuntjw/go-gin-starter/configs"
	"github.com/nuntjw/go-gin-starter/database"
	"github.com/nuntjw/go-gin-starter/server"
)

func main() {
	configs.LoadEnv()
	database.ConnectMongoDB()
	server.InitRoute()
	server.Run()
}
