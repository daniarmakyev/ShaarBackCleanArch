// route/user_router.go
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

func NewUserRouter(env *bootstrap.Env, group *gin.RouterGroup, timeout time.Duration, db *sql.DB) {
	ur := repository.NewUserRepository(db)

	uc := &controller.UserController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		UserUsecase:         usecase.NewUserUsecase(ur),
		Env:                 env,
	}

	group.GET("/user", uc.GetUser)
	group.PATCH("/user", uc.UpdateUser)
}
