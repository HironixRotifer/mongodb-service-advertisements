package main

import (
	"log"
	"os"

	// "github.com/HironixRotifer/mongodb-service-advertisements/controllers"
	// "github.com/HironixRotifer/mongodb-service-advertisements/database"
	"github.com/HironixRotifer/mongodb-service-advertisements/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.Routes(router)

	router.GET("")

	log.Fatal(router.Run(":" + port))

}
