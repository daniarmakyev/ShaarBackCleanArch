package route

import (
	"database/sql"
	"time"

	"shaar/api/controller"
	"shaar/bootstrap"
	"shaar/repository"
	"shaar/usecase"

	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRouter(env *bootstrap.Env, group *gin.RouterGroup, timeout time.Duration, db *sql.DB) {
	ur := repository.NewUserRepository(db)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
