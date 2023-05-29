package main

import (
	"net/http"
	"weather/internal"

	"github.com/gin-gonic/gin"
)

// UpdateWeather is a Gin handler that handles weather update requests.
// It extracts the weather value from the request body, checks if the address is
// registered in the contract then adds it to the WeatherUpdate channel,
// and returns a response indicating successful processing.
func UpdateWeather(wu *internal.WeatherUpdate) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the weather from the body
		var floatValue float64
		if err := c.ShouldBindJSON(&floatValue); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//TODO: Check to see if the address is registered in the smart contract or not.

		// Pass it to a channel
		wu.AddToChannel(floatValue)

		c.JSON(http.StatusOK, gin.H{"message": "Float received and added to the processing queue"})
	}
}
