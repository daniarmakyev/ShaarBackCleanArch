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
	protectedRouter := router.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	protectedRouter.GET("/profile", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This is the profile route",
		})
	})
}
