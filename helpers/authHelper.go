package helpers

import (
	"context"
	"errors"

	models "github.com/HousewareHQ/backend-engineering-octernship/internal/app/_your_app_/models"
	"go.mongodb.org/mongo-driver/bson"
)

// var userTokenCollection *mongo.Collection = database.OpenCollection(database.Client, "usertokens")

// func CheckUserType(c *gin.Context, role string)(err error){
// 	userType := c.GetString("user_type")
// 	err = nil

// 	if userType != role {
// 		err = errors.New("Unauthorized to access this resource")
// 		return err
// 	}
// 	return err
// }

// func MatchUserTypeToUid(c *gin.Context, userId string)(err error){
// 	userType := c.GetString("user_type")
// 	uid := c.GetString("uid")
// 	err = nil

// 	if userType == "USER" && uid != userId {
// 		err = errors.New("Unauthorized to access this resource")
// 		return err
// 	}
// 	err = CheckUserType(c, userType)
// 	return err
// }

func MatchClientTokenToUserToken(userid string, clientToken string) error {
	var user models.UserTokens
	filter := bson.M{"userid": userid}
	err := userTokenCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return err
	}
	if *user.Token != clientToken{
		return errors.New("user session logged out, you need to login")
	}
	return nil
}
