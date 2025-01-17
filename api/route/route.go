package route

import (
	"database/sql"
	"shaar/bootstrap"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, db *sql.DB, router *gin.Engine) {
	publicRouter := router.Group("")
	NewSignupRouter(publicRouter, db)
	NewSigninRouter(env, publicRouter, db)
}
