package main

import (
	"log"
	"os"

	routes "github.com/HousewareHQ/backend-engineering-octernship/internal/app/auth/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	// if no port is specified, use port 8000
	if port==""{
		port="8000"
	}

	router := gin.New() // create a new gin
	router.Use(gin.Logger()) // use gin logger for the fancy logs

	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.AdminRoutes(router)
	
	// this is a public route
  	router.GET("/api-1", func(c *gin.Context){
    	c.JSON(200, gin.H{"success": "Access granted for api-1"})
  	})
	
	router.Run(":" + port)
}
