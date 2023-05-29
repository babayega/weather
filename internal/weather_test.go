package internal_test

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"weather/internal"
)

func TestWeatherUpdate(t *testing.T) {
	// Create a new instance of WeatherUpdate
	w := internal.NewWeatherUpdate()
	go w.Update()

	// Add a value to the channel
	w.AddToChannel(25.5)

	// Wait for the channel processing to complete
	time.Sleep(100 * time.Millisecond)

	// Get the latest weather value
	latestWeather := w.GetLatestWeather()

	// Assert that the latest weather value is updated correctly
	require.Equal(t, 25.5, latestWeather, "Latest weather value should be 25.5")
}

func TestWeatherUpdate_RaceCondition(t *testing.T) {
	// Create a new instance of WeatherUpdate
	w := internal.NewWeatherUpdate()
	go w.Update()

	// Set up a wait group to ensure all goroutines are completed
	var wg sync.WaitGroup
	wg.Add(2)

	// Simulate multiple goroutines adding values to the channel concurrently
	go func() {
		defer wg.Done()
		w.AddToChannel(20.0)
	}()

	go func() {
		defer wg.Done()
		w.AddToChannel(22.5)
	}()

	// Wait for the goroutines to complete
	wg.Wait()

	// Wait for the channel processing to complete
	time.Sleep(100 * time.Millisecond)

	// Get the latest weather value
	latestWeather := w.GetLatestWeather()

	// Assert that the latest weather value is one of the added values
	require.Contains(t, []float64{20.0, 22.5}, latestWeather, "Latest weather value should be one of the added values")
}
