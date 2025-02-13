package domain

import "context"

type LocationAir struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type PollutionAir struct {
	Timestamp string `json:"ts"`
	AQIUS     int    `json:"aqius"`
	MainUS    string `json:"mainus"`
	AQICN     int    `json:"aqicn"`
	MainCN    string `json:"maincn"`
}

type WeatherAir struct {
	Timestamp     string  `json:"ts"`
	Temperature   int     `json:"tp"`
	Pressure      int     `json:"pr"`
	Humidity      int     `json:"hu"`
	WindSpeed     float64 `json:"ws"`
	WindDirection int     `json:"wd"`
	Icon          string  `json:"ic"`
}

type CurrentAir struct {
	Pollution PollutionAir `json:"pollution"`
	Weather   WeatherAir   `json:"weather"`
}

type DataAir struct {
	City     string      `json:"city"`
	State    string      `json:"state"`
	Country  string      `json:"country"`
	Location LocationAir `json:"location"`
	Current  CurrentAir  `json:"current"`
}

type AirResponse struct {
	Status string  `json:"status"`
	Data   DataAir `json:"data"`
}

type SimpleAir struct {
	City  string `json:"city"`
	AQIUS int    `json:"aqius"`
}

type AirUsecase interface {
	GetAir(ctx context.Context) (error, *SimpleAir)
}
