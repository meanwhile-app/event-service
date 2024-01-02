package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/nuntjw/go-gin-starter/controllers"
)

func InitUserRoute(routerGroup *gin.RouterGroup) {
	ctrl := controllers.NewEventController()

	userGroup := routerGroup.Group("/events")
	{
		userGroup.GET("/", ctrl.GetEvents)
	}
}
