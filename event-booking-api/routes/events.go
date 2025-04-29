package routes

import (
	"github.com/Lucas-Mol/go-studies/event-booking-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func createEvent(c *gin.Context) {
	userID := c.GetInt64("userId")

	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	event.UserID = userID
	err = event.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create an event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created",
		"event":   event,
	})
}

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
		return
	}
	c.JSON(http.StatusOK, events)
}

func getEventById(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch an event"})
		return
	}

	c.JSON(http.StatusOK, event)
}

func updateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch an event"})
		return
	}

	if event.UserID != c.GetInt64("userId") {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized to update this event"})
		return
	}

	var updatedEvent models.Event
	err = c.ShouldBindJSON(&updatedEvent)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated",
		"event":   updatedEvent,
	})
}

func deleteEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch an event"})
		return
	}

	if event.UserID != c.GetInt64("userId") {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized to delete this event"})
		return
	}

	err = event.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted"})
}
