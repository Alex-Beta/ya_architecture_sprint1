package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// DeviceService handles fetching device data from external API
type DeviceService struct {
	BaseURL    string
	HTTPClient *http.Client
}

// SensorResponse represents the response from the device API
type SensorResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorType  string    `json:"type"`
	LastUpdated time.Time `json:"last_updated"`
	CreatedAt   time.Time `json:"created_at"`
}

// CreateSensorRequest представляет данные для создания сенсора во внешнем API
type CreateSensorRequest struct {
	Type     string `json:"type"`
	Location string `json:"location"`
	Name     string `json:"name"`
	Unit     string `json:"unit"`
}

// GetSensorsResponse представляет ответ со списком сенсоров от API
type GetSensorsResponse struct {
	Sensors []SensorResponse `json:"sensors"`
}

// NewDeviceService creates a new device service
func NewDeviceService(baseURL string) *DeviceService {
	return &DeviceService{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetSensor получает сенсор по идентификатору
func (s *DeviceService) GetSensor(id int) (*SensorResponse, error) {
	// Формируем URL для запроса
	url := fmt.Sprintf("%s/api/sensor/%d", s.BaseURL, id)

	// Создаем GET-запрос
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %w", err)
	}

	// Выполняем запрос
	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка отправки запроса: %w", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("внешнее API вернуло неожиданный статус: %d, тело: %s",
			resp.StatusCode, string(body))
	}

	// Декодируем ответ напрямую в массив
	var sensor SensorResponse
	if err := json.NewDecoder(resp.Body).Decode(&sensor); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	return &sensor, nil
}
