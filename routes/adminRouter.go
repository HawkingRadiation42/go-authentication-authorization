package routes

import (
	controller "github.com/HousewareHQ/backend-engineering-octernship/internal/app/_your_app_/controllers"
	"github.com/HousewareHQ/backend-engineering-octernship/internal/app/_your_app_/middleware"

	"github.com/gin-gonic/gin"
)


func AdminRoutes(incomingRoutes *gin.Engine) {
	// only admin usertype can access these request - using the authenticate middleware to check user type
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("/users", controller.AddUserToOrg())
	incomingRoutes.DELETE("/users/{userId}", controller.DeleteUserFromOrg())
}
	// this orgId should match the orgId of the ADMIN - orgId is in the token
	// Admin User adds a new User account(by providing the username & password)
	
	// this orgId should match the orgId of the ADMIN and only admin can send this request  - middleware will check this
	// we have to delete the user from the database and also from the organization


