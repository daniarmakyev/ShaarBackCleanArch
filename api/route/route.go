package route

import (
	"database/sql"
	"shaar/bootstrap"
	"shaar/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db *sql.DB, router *gin.Engine) {
	publicRouter := router.Group("")
	NewSignupRouter(env, publicRouter, timeout, db)
	NewSigninRouter(env, publicRouter, timeout, db)
	NewRefreshTokenRouter(env, publicRouter, timeout, db)
	NewEventRouter(env, publicRouter, timeout, db)
	protectedRouter := router.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	NewWeatherRouter(env, protectedRouter, timeout)
	NewAirRouter(env, protectedRouter, timeout)
	NewPlacesRouter(env, protectedRouter, timeout, db)
	NewUserRouter(env, protectedRouter, timeout, db)
}
