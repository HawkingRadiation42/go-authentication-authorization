package routes

import (
	controller "github.com/HousewareHQ/backend-engineering-octernship/internal/app/auth/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("login", controller.Login())	
	incomingRoutes.POST("signup", controller.Signup()) // added only for adding admin data to the database
}
