package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/meanwhile-app/event-service/controllers"
	"github.com/meanwhile-app/event-service/middewares"
)

func InitUserRoute(routerGroup *gin.RouterGroup) {
	ctrl := controllers.NewEventController()

	eventGroup := routerGroup.Group("/events", middewares.Authorize())
	{
		eventGroup.GET("/", ctrl.GetEvents)
		eventGroup.POST("/", ctrl.InsertEvent)
		eventGroup.GET("/nearby", ctrl.GetNearbyEvents)
	}
}
