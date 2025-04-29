package routes

import (
	"github.com/Lucas-Mol/go-studies/event-booking-api/models"
	"github.com/Lucas-Mol/go-studies/event-booking-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signup(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": struct {
			ID    int64  `json:"id"`
			Email string `json:"email"`
		}{
			ID:    user.ID,
			Email: user.Email,
		},
	})
}

func login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not authenticate user",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Could not authenticate user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
