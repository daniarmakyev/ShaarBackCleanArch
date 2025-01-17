package main

import (
	"log"
	"shaar/api/route"
	"shaar/bootstrap"
	"shaar/postgres"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	db, dbError := postgres.ConnectDB()

	if dbError != nil {
		log.Fatalf("Error connecting to database: %v", dbError)
	}
	defer postgres.CloseDB(db)

	router := gin.Default()

	env := bootstrap.NewEnv()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	route.Setup(env, db, router)

	err := router.Run(":8080")

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
