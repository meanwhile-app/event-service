package main

import (
	"github.com/nuntjw/go-gin-starter/configs"
	"github.com/nuntjw/go-gin-starter/server"
)

func main() {
	configs.LoadEnv()
	// server.ConnectDB()
	server.InitRoute()
	server.Run()
}
