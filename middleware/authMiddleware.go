package middleware

import (
	"fmt"
	"net/http"

	helper "github.com/HousewareHQ/backend-engineering-octernship/internal/app/_your_app_/helpers"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc{
	return func(c *gin.Context){
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			c.Abort()
			return 
		}
		claims, err := helper.ValidateToken(clientToken)
		if err != ""{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return 
		}
		if claims.UserType != "ADMIN"{
			c.JSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("Unauthorized to access this resource")})
			c.Abort()
			return 
		}
		c.Set("username", claims.Username)
		c.Set("uid", claims.Uid)
		// c.Set("orgid", claims.Orgid) // dont know if its right or not
		c.Next()
	}
}

func TokenAuth() gin.HandlerFunc{
	return func(c *gin.Context){
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			c.Abort()
			return 
		}
		claims, err := helper.ValidateToken(clientToken)
		if err != ""{
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return 
		}
		errr := helper.MatchClientTokenToUserToken(claims.Uid, clientToken)
		if errr != nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user session logged out, you need to login"})
			c.Abort()
			return
		}
		c.Set("orgId", claims.Orgid)
		c.Next()
	}
}