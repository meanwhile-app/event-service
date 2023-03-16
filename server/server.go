package server

import (
	"github.com/gin-gonic/gin"
	"github.com/nuntjw/go-gin-starter/configs"
	"github.com/nuntjw/go-gin-starter/server/routing"
)

var r *gin.Engine

func InitRoute() {
	r = gin.Default()

	apiV1Group := r.Group("/api/v1")
	{
		routing.InitUserRoute(apiV1Group)
	}

}

func Run() {
	env := configs.GetEnv()
	port := ":" + env["PORT"]
	r.Run(port)
}
