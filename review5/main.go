package main

import (
	"fmt"

	"KATAreview/review5/weather"
)

func main() {
	weatherData := []byte(`{"temperature":"+5 °C","wind":"7 km/h","description":"Light rain","forecast":[{"day":"1","temperature":"1 °C","wind":"15 km/h"},{"day":"2","temperature":"+4 °C","wind":"6 km/h"},{"day":"3","temperature":" °C","wind":"0 km/h"}]}`)
	fmt.Println(weather.GetWeather(weatherData))
}
