package middewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/meanwhile-app/event-service/utilities"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")

		if authorization == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   "Authorization is required",
			})
			c.Abort()
			return
		}

		token := strings.Split(authorization, " ")[1]

		claims, err := utilities.DecodeJwt(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}

		if uid, exists := claims["uid"].(string); exists {
			c.Set("user_id", uid)
		}

		c.Next()
	}
}
