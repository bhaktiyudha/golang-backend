package main

import (
	"log"
	"mission-2/config"
	"mission-2/controller"
	"mission-2/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(middleware.CORS)
	router.Use(gin.LoggerWithFormatter(middleware.LoggerMiddleware))
	countGroup := router.Group("/movie")
	countGroup.GET("", controller.GetMovie)
	countGroup.GET("/:id", controller.GetMovieDetail)

	err := router.Run(config.APIPort)

	if err != nil {
		log.Fatalf("Error run API : %s", err)
	}
	log.Printf("Running")
}
