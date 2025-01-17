package route

import (
	"database/sql"
	"shaar/api/controller"
	"shaar/repository"
	"shaar/usecase"

	"github.com/gin-gonic/gin"
)

func NewSignupRouter(group *gin.RouterGroup, db *sql.DB) {
	ur := repository.NewUserRepository(db)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur),
	}
	group.POST("/signup", sc.Signup)
}
