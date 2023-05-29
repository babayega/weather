package main

import (
	"weather/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Middlewares
	router.Use(VerifySignatureMiddleware(), WeatherRateLimitMiddleware())

	// Start the weather update service
	wu := internal.NewWeatherUpdate()
	go wu.Update()

	// Route handler for the POST /process endpoint
	router.POST("/process", UpdateWeather(wu))

	// Run the server on port 8080
	router.Run(":8080")
}
