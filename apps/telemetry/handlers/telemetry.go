package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"telemetry/services"
	"time"
)

func MakeTelemetryHandler(deviceService *services.DeviceService, temperatureService *services.TemperatureService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg, err := handleTelemetry(w, r, deviceService, temperatureService)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(msg)
	}
}

type TelemetryMessage struct {
	Value     float64   `json:"value"`
	Unit      string    `json:"unit"`
	Timestamp time.Time `json:"timestamp"`
	SensorID  string    `json:"sensor_id"`
	Location  string    `json:"location"`
}

// Проверяет создан ли сенсор у апи девайсов, запрашивает температуру у апи температуры, и возвращает её в случае успеха
func handleTelemetry(w http.ResponseWriter, r *http.Request, ds *services.DeviceService, ts *services.TemperatureService) (*TelemetryMessage, error) {
	parts := strings.Split(r.URL.Path, "/")
	sensorID, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, fmt.Errorf("неверный формат ID сенсора: %w", err)
	}

	sensor, err := ds.GetSensor(sensorID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return nil, err
	}

	temp, err := ts.GetTemperatureByID(parts[2])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return nil, err
	}

	response := TelemetryMessage{
		SensorID:  parts[2],
		Value:     temp.Value,
		Unit:      sensor.Unit,
		Timestamp: temp.Timestamp,
	}

	return &response, nil
}
