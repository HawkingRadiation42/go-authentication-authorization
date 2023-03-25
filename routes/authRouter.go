package routes

import (
	controller "github.com/HousewareHQ/backend-engineering-octernship/internal/app/_your_app_/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("login", controller.Login())	
	incomingRoutes.POST("signup", controller.Signup()) // added only for adding admin data to the database
}
