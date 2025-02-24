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

func NewEventRouter(env *bootstrap.Env, group *gin.RouterGroup, timeout time.Duration, db *sql.DB) {
	er := repository.NewEventRepository(db)
	ec := controller.NewEventController(
		usecase.NewEventUsecase(er, timeout),
	)

	group.GET("/events", ec.GetAllEvents)
	group.POST("/events", ec.Create)
}
