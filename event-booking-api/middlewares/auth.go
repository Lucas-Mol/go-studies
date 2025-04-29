package middlewares

import (
	"github.com/Lucas-Mol/go-studies/event-booking-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	userID, err := utils.ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	c.Set("userId", userID)
	c.Next()
}
