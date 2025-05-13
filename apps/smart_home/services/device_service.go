package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"smarthome/models"
	"time"
)

// DeviceService handles fetching device data from external API
type DeviceService struct {
	BaseURL    string
	HTTPClient *http.Client
}

// CreateSensorResponse represents the response from the device API
type CreateSensorResponse struct {
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
	Sensors []CreateSensorResponse `json:"sensors"`
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

// CreateSensor отправляет запрос на создание сенсора во внешнее API
func (s *DeviceService) CreateSensor(sensor *models.SensorCreate) (*CreateSensorResponse, error) {
	// Формируем запрос для внешнего API
	createRequest := CreateSensorRequest{
		Type:     string(sensor.Type),
		Location: sensor.Location,
		Name:     sensor.Name,
		Unit:     sensor.Unit,
	}

	// Сериализуем данные в JSON
	jsonData, err := json.Marshal(createRequest)
	if err != nil {
		return nil, fmt.Errorf("ошибка сериализации данных сенсора: %w", err)
	}

	// Формируем URL для запроса
	url := fmt.Sprintf("%s/api/sensor", s.BaseURL)

	// Отправляем POST-запрос
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("ошибка создания запроса: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Выполняем запрос
	resp, err := s.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка отправки запроса: %w", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("внешнее API вернуло неожиданный статус: %d, тело: %s",
			resp.StatusCode, string(body))
	}

	// Декодируем ответ
	var deviceResp CreateSensorResponse
	if err := json.NewDecoder(resp.Body).Decode(&deviceResp); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	return &deviceResp, nil
}

// GetSensors получает список всех сенсоров из внешнего API
func (s *DeviceService) GetSensors() ([]CreateSensorResponse, error) {
	// Формируем URL для запроса
	url := fmt.Sprintf("%s/api/sensor/GetAllSensors", s.BaseURL)

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
	var sensors []CreateSensorResponse
	if err := json.NewDecoder(resp.Body).Decode(&sensors); err != nil {
		return nil, fmt.Errorf("ошибка декодирования ответа: %w", err)
	}

	return sensors, nil
}
