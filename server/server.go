package server

import (
	"github.com/gin-gonic/gin"
	"github.com/meanwhile-app/event-service/configs"
	"github.com/meanwhile-app/event-service/server/routing"
)

var r *gin.Engine

func InitRoute() {
	r = gin.Default()

	apiV1Group := r.Group("/api/v1")
	{
		routing.InitEventRoute(apiV1Group)
	}

}

func Run() {
	env := configs.GetEnv()
	port := ":" + env["PORT"]
	r.Run(port)
}
