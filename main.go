package main

import (
	"log"
	"os"

	routes "github.com/HousewareHQ/backend-engineering-octernship/internal/app/_your_app_/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port==""{
		port="8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.AdminRoutes(router)

  	router.GET("/api-1", func(c *gin.Context){
    	c.JSON(200, gin.H{"success": "Access granted for api-1"})
  	})
	
	router.Run(":" + port)
}
