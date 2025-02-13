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

type airUsecase struct {
	contextTimeout time.Duration
}

func NewAirUsecase(timeout time.Duration) domain.AirUsecase {
	return &airUsecase{
		contextTimeout: timeout,
	}
}

func (wu *airUsecase) GetAir(ctx context.Context) (error, *domain.SimpleAir) {
	ctx, cancel := context.WithTimeout(ctx, wu.contextTimeout)
	defer cancel()
	apiKey := os.Getenv("AIR_API_KEY")
	if apiKey == "" {
		return fmt.Errorf("API key not found"), nil
	}

	req, err := http.NewRequestWithContext(ctx, "GET", "http://api.airvisual.com/v2/city?city=Bishkek&state=Bishkek&country=Kyrgyzstan&key="+apiKey, nil)
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
		return fmt.Errorf("failed to get air data: %s", resp.Status), nil
	}

	var airResponse domain.AirResponse
	if err := json.NewDecoder(resp.Body).Decode(&airResponse); err != nil {
		return err, nil
	}

	simpleAir := domain.SimpleAir{
		City:  airResponse.Data.City,
		AQIUS: airResponse.Data.Current.Pollution.AQIUS,
	}
	return nil, &simpleAir
}
