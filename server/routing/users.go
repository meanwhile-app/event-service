package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/nuntjw/go-gin-starter/controllers"
)

func InitUserRoute(routerGroup *gin.RouterGroup) {
	user := controllers.NewUserController()

	userGroup := routerGroup.Group("/users")
	{
		userGroup.GET("/", user.GetUsers)
	}
}
