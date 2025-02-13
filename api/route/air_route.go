package route

import (
	"shaar/api/controller"
	"shaar/bootstrap"
	"shaar/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewAirRouter(env *bootstrap.Env, group *gin.RouterGroup, timeout time.Duration) {
	rtc := controller.NewAirController(
		usecase.NewAirUsecase(timeout),
	)
	group.GET("/air", rtc.GetAir)
}
