package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"shaar/domain"
	"time"
)

type weatherUsecase struct {
	contextTimeout time.Duration
}

func NewWeatherUsecase(timeout time.Duration) domain.WeatherUsecase {
	return &weatherUsecase{
		contextTimeout: timeout,
	}
}

func (wu *weatherUsecase) GetWeather(ctx context.Context) (error, *domain.SimpleWeather) {
	ctx, cancel := context.WithTimeout(ctx, wu.contextTimeout)
	defer cancel()
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("API key not found"), nil
	}

	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.weatherapi.com/v1/current.json?q=Bishkek&lang=en&key="+apiKey, nil)
	if err != nil {
		return err, nil
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if os.IsTimeout(err) {
			return fmt.Errorf("request timed out"), nil
		}
		return err, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get weather data: %s", resp.Status), nil
	}

	var weatherResponse domain.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return err, nil
	}

	simpleWeather := domain.SimpleWeather{
		Region: weatherResponse.Location.Region,
		TempC:  weatherResponse.Current.TempC,
	}
	return nil, &simpleWeather
}
