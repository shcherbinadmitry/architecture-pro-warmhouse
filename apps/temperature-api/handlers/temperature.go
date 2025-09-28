package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"math/rand/v2"
	"net/http"
	"time"
)

// TemperatureResponse represents the response for API
type TemperatureResponse struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

// TemperatureHandler handles temperature-related requests
type TemperatureHandler struct{}

// NewTemperatureHandler creates a new TemperatureHandler
func NewTemperatureHandler() *TemperatureHandler {
	return &TemperatureHandler{}
}

// RegisterRoutes registers the temperature routes
func (h *TemperatureHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/temperature", h.GetTemperatureByLocation)
	router.GET("/temperature/:id", h.GetTemperatureById)

}

// GetTemperatureById handles GET /temperature/:id
func (h *TemperatureHandler) GetTemperatureById(c *gin.Context) {
	id := c.Param("id")
	location, id := getLocationAndSensorId("", id)
	// Create temperature data for response
	response := TemperatureResponse{
		Value:       getRandomTemperature(),
		Unit:        "Celsius",
		Timestamp:   time.Now(),
		Location:    location,
		Status:      "active",
		SensorType:  "temperature",
		SensorID:    id,
		Description: fmt.Sprintf("Temperature for sensor with id: %s", id),
	}

	// Return the temperature data
	c.JSON(http.StatusOK, response)
}

// GetTemperatureByLocation handles GET /temperature
func (h *TemperatureHandler) GetTemperatureByLocation(c *gin.Context) {
	// Get location from query parameter
	location := c.DefaultQuery("location", "")
	location, id := getLocationAndSensorId(location, "")
	// Create temperature data for response
	response := TemperatureResponse{
		Value:       getRandomTemperature(),
		Unit:        "°C",
		Timestamp:   time.Now(),
		Location:    location,
		Status:      "active",
		SensorType:  "temperature",
		SensorID:    id,
		Description: fmt.Sprintf("Temperature in %s", location),
	}

	// Return the temperature data
	c.JSON(http.StatusOK, response)
}

// getRandomTemperature returns a random temperature between -35 and 75
func getRandomTemperature() float64 {
	const minTemp = -35.0
	const maxTemp = 75.0
	val := minTemp + rand.Float64()*(maxTemp-minTemp)
	return math.Round(val*100) / 100
}

// getLocationAndSensorId returns a location and sensor ID based on the provided location and sensor ID
func getLocationAndSensorId(location string, sensorID string) (string, string) {
	// If no location is provided, use a default based on sensor ID
	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	// If no sensor ID is provided, generate one based on location
	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}

	return location, sensorID
}
