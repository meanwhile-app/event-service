package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nuntjw/go-gin-starter/models"
	"github.com/nuntjw/go-gin-starter/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		c.Abort()
		return
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
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    events,
	})
}

func (ctrl *EventController) InsertEvent(c *gin.Context) {
	var reqBody types.InsertEventPayload
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := primitive.ObjectIDFromHex("65941000f75468983b3903c8")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "insert error",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	payload := types.InsertEventPayload{
		Title:     reqBody.Title,
		Location:  reqBody.Location,
		CreatedBy: userId,
	}

	result, err := ctrl.eventModel.InsertOne(&payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "insert error",
			"error":   err.Error(),
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"data":    result,
	})
}
