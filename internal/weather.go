package internal

import "fmt"

// WeatherUpdate represents a weather update and provides functionality to process and store weather data.
type WeatherUpdate struct {
	wChannel      chan float64 // wChannel is a channel used to receive and process float64 weather values.
	latestWeather float64      // latestWeather holds the latest processed weather value.
}

// NewWeatherUpdate creates a new instance of WeatherUpdate and initializes the wChannel.
func NewWeatherUpdate() *WeatherUpdate {
	return &WeatherUpdate{wChannel: make(chan float64), latestWeather: 0}
}

// Update continuously processes the float values received from the wChannel.
// It updates the latestWeather field.
func (w *WeatherUpdate) Update() {
	for floatValue := range w.wChannel {
		fmt.Printf("Processing float: %.2f\n", floatValue)

		// Update the value of latestWeather
		w.latestWeather = floatValue
	}
}

// AddToChannel adds a float value to the wChannel for processing by the Update method.
func (w *WeatherUpdate) AddToChannel(wVal float64) {
	w.wChannel <- wVal
}

// GetLatestWeather returns the latest processed weather value.
func (w WeatherUpdate) GetLatestWeather() float64 {
	return w.latestWeather
}
