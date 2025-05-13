package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

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

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
    parts := strings.Split(r.URL.Path, "/")
	location := r.URL.Query().Get("location")
	sensorID := parts[2]
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

	rand.Seed(time.Now().UnixNano())
	temp := rand.Float64()*60 - 10

	response := TemperatureResponse{
		Location:    location,
		SensorID:    sensorID,
		Value: temp,
		Unit:        "°C",
		Timestamp:   time.Now(),
	}

    // Логирование ответа
	responseJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Printf("Error marshaling response: %v", err)
	} else {
		log.Printf("Response to %s:\n%s", r.URL.Path, responseJSON)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/temperature/", temperatureHandler)
	log.Println("Server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
