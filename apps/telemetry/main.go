package main

import (
	"log"
	"net/http"
	"os"
	"telemetry/handlers"
	"telemetry/services"
)

func main() {

	// Initialize temperature service
	temperatureAPIURL := getEnv("TEMPERATURE_API_URL", "http://temperature-api:8081")
	temperatureService := services.NewTemperatureService(temperatureAPIURL)
	log.Printf("Temperature service initialized with API URL: %s\n", temperatureAPIURL)

	// Initialize device service
	deviceAPIURL := getEnv("DEVICE_API_URL", "http://devices-api:8080")
	deviceService := services.NewDeviceService(deviceAPIURL)
	log.Printf("Device service initialized with API URL: %s\n", deviceAPIURL)

	// Регистрируем обработчик с передачей deviceService
	http.HandleFunc("/telemetry/", handlers.MakeTelemetryHandler(deviceService, temperatureService))

	log.Println("Server started on :8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
