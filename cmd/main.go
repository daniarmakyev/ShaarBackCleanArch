package main

import (
	"log"
	"shaar/api/route"
	"shaar/bootstrap"
	"shaar/postgres"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	db, dbError := postgres.ConnectDB()

	if dbError != nil {
		log.Fatalf("Error connecting to database: %v", dbError)
	}
	defer postgres.CloseDB(db)

	env := bootstrap.NewEnv()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	route.Setup(env, timeout, db, router)

	err := router.Run(":8080")

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
