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

func NewPlacesRouter(env *bootstrap.Env, group *gin.RouterGroup, timeout time.Duration, db *sql.DB) {
	pr := repository.NewPlacesRepository(db)
	pc := controller.PlacesController{
		PlacesUsecase: usecase.NewPlacesUsecase(pr, timeout),
	}

	group.GET("/places", pc.GetAllPlaces)
	group.GET("/categories/:categories", pc.GetPlacesByCategory)
	group.GET("/places/price/:price", pc.GetPlacesByPrice)
	group.GET("/categories", pc.GetAllCategories)
}
