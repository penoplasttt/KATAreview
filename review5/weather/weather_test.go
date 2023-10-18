package weather

import (
	"reflect"
	"testing"
	//mock "github.com/penoplasttt/review/review5/weather/mocks"
)

func TestGetWeather(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Weather
		wantErr bool
	}{
		{
			name: "ok",
			args: args{data: []byte(`{"temperature":"+5 °C","wind":"7 km/h","description":"Light rain","forecast":[{"day":"1","temperature":"1 °C","wind":"15 km/h"}]}`)},
			want: Weather{Temperature: "+5 °C", Wind: "7 km/h", Description: "Light rain", Forecast: []Forecast{{Day: "1", Temperature: "1 °C", Wind:"15 km/h" }}},
			wantErr: false,

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetWeather(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWeather() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWeather() = %v, want %v", got, tt.want)
			}
		})
	}
}
