package main

import "fmt"

type WeatherService interface {
    GetWeather() (Weather, error)
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

func GetWeather(ws WeatherService) {
    data, err := ws.GetWeather()
    if err != nil {
        fmt.Printf("%v", err)
        return
    }
    fmt.Println(data.Temperature)
    fmt.Println(data.Wind)
    fmt.Println(data.Description)

    for _, day := range data.Forecast {
        fmt.Println(day.Day, day.Wind, day.Temperature)
    }
}

type MockService struct{}

func (m *MockService) GetWeather() (Weather, error) {
    data := Weather{
        Temperature: "+3 °C",
        Wind:        "31 km/h",
        Description: "Rain and snow shower",
        Forecast: []Forecast{
            {Day: "1", Temperature: "+4 °C", Wind: "32 km/h"},
            {Day: "2", Temperature: "+6 °C", Wind: "15 km/h"},
            {Day: "3", Temperature: "+5 °C", Wind: "13 km/h"},
        },
    }
    return data, nil
}

func main() {
    weather := &MockService{}
    GetWeather(weather)
}