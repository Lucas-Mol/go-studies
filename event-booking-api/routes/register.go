package routes

import (
	"github.com/Lucas-Mol/go-studies/event-booking-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered on event",
	})
}

func cancelRegistration(c *gin.Context) {
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find event"})
		return
	}

	err = event.CancelRegistration(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User cancelled registration"})
}
