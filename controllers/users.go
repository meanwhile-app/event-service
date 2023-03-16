package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nuntjw/go-gin-starter/models"
)

type UserController struct {
	userModel models.UserModel
}

func NewUserController() *UserController {
	return &UserController{
		userModel: models.UserModel{},
	}
}

func (controller *UserController) GetUsers(c *gin.Context) {
	users, err := controller.userModel.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "get user error.",
			"error":   err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    users,
	})
}
