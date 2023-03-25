package routes

import (
	controller "github.com/HousewareHQ/backend-engineering-octernship/internal/app/_your_app_/controllers"
	middleware "github.com/HousewareHQ/backend-engineering-octernship/internal/app/_your_app_/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.TokenAuth())
	incomingRoutes.GET("users", controller.GetUsers())
	incomingRoutes.POST("logout", controller.Logout())
	// this orgId should match that of the user/admin, then only can send this request
}


