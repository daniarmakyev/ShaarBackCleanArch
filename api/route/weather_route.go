package route

import (
	"shaar/api/controller"
	"shaar/bootstrap"
	"shaar/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewWeatherRouter(env *bootstrap.Env, group *gin.RouterGroup, timeout time.Duration) {
	rtc := controller.NewWeatherController(
		usecase.NewWeatherUsecase(timeout),
	)
	group.GET("/weather", rtc.GetWeather)
}
