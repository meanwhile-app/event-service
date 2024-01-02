package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/nuntjw/go-gin-starter/controllers"
	"github.com/nuntjw/go-gin-starter/middewares"
)

func InitUserRoute(routerGroup *gin.RouterGroup) {
	ctrl := controllers.NewEventController()

	eventGroup := routerGroup.Group("/events", middewares.Authorize())
	{
		eventGroup.GET("/", ctrl.GetEvents)
		eventGroup.GET("/nearby", ctrl.GetNearbyEvents)
	}
}
