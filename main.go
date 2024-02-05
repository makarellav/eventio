package main

import (
	"github.com/gin-gonic/gin"
	"github.com/makarellav/eventio/models"
	"log"
	"net/http"
)

var id int

func main() {
	r := gin.Default()

	r.GET("/events", getAllEvents)
	r.POST("/events", createEvent)

	log.Fatal(r.Run())
}

func createEvent(c *gin.Context) {
	id++

	var event models.Event
	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not parse the request data",
		})

		return
	}

	event.ID = id
	event.UserID = id
	event.Save()

	c.JSON(http.StatusCreated, gin.H{"event": event})
}

func getAllEvents(c *gin.Context) {
	events := models.GetAllEvents()

	c.JSON(http.StatusOK, events)
}
