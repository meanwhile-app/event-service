package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nuntjw/go-gin-starter/models"
	"go.mongodb.org/mongo-driver/bson"
)

type EventController struct {
	eventModel *models.EventModel
}

func NewEventController() *EventController {
	return &EventController{
		eventModel: models.NewEventModel(),
	}
}

func (ctrl *EventController) GetEvents(c *gin.Context) {
	events, err := ctrl.eventModel.GetEvents(bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "get user error.",
			"error":   err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    events,
	})
}

func (ctrl *EventController) GetNearbyEvents(c *gin.Context) {
	location := strings.Split(c.Query("location"), ",")
	events, err := ctrl.eventModel.GetNearbyEvents(location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "get user error.",
			"error":   err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    events,
	})
}
