package controller

import (
	"net/http"
	"shaar/domain"

	"github.com/gin-gonic/gin"
)

type WeatherController struct {
	WeatherUsecase domain.WeatherUsecase
}

func NewWeatherController(weatherUsecase domain.WeatherUsecase) *WeatherController {
	return &WeatherController{
		WeatherUsecase: weatherUsecase,
	}
}

func (wc *WeatherController) GetWeather(c *gin.Context) {
	err, weatherResponse := wc.WeatherUsecase.GetWeather(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, weatherResponse)
}
