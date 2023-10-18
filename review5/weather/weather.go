package weather

import "encoding/json"


//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name=WeatherService
type WeatherService interface {
	GetWeather(data []byte) (Weather, error)
}

type Weather struct {
	Temperature string     `json:"temperature"`
	Wind        string     `json:"wind"`
	Description string     `json:"description"`
	Forecast    []Forecast `json:"forecast"`
}

type Forecast struct {
	Day         string `json:"day"`
	Temperature string `json:"temperature"`
	Wind        string `json:"wind"`
}

func GetWeather(data []byte) (Weather, error) {
	var weather Weather
	err := json.Unmarshal(data, &weather)
	if err != nil {
		return weather, err
	}
	return weather, err
}
