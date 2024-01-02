package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nuntjw/go-gin-starter/models"
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
	events, err := ctrl.eventModel.GetEvents()
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
