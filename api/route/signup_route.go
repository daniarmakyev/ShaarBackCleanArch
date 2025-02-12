package route

import (
	"database/sql"
	"shaar/api/controller"
	"shaar/bootstrap"
	"shaar/repository"
	"shaar/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func NewSignupRouter(env *bootstrap.Env, group *gin.RouterGroup, timeout time.Duration, db *sql.DB) {
	ur := repository.NewUserRepository(db)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
