package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/HousewareHQ/backend-engineering-octernship/internal/app/_your_app_/database"

	jwt "github.com/dgrijalva/jwt-go"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type SignedDetails struct {
	Username string
	Uid string
	Orgid string
	UserType string
	jwt.StandardClaims
}

var userTokenCollection *mongo.Collection = database.OpenCollection(database.Client, "usertokens")
var SECRET_KEY string = os.Getenv("SECRET_KEY")

// function input username, uid, orgid
func GenerateAllTokens(username string, uid string, orgid string, usertype string)(signedToken string, signedRefreshToken string, err error){
	claims := &SignedDetails{
		Username: username,
		Uid: uid,
		Orgid: orgid,
		UserType: usertype,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour).Unix(),
		},
	}
	refreshClaims := &SignedDetails{
		Username: username,
		Uid: uid,
		Orgid: orgid,
		UserType: usertype,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	
	if err != nil {
		log.Panic(err)
		return 
	}
	return token, refreshToken, err

}
//TODO: add refresh Access token function code where refresh token is accepted and access token is generated and saved to the database and accesstoken returned
// func RefreshAccessToken(signedToken string)


func ValidateToken(signedToken string) (claims *SignedDetails, msg string){
	// println(signedToken) //debugging
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{}, 
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		// msg = err.Error() // debugging
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("the token is expired")
		msg = err.Error()
		return
	}
	return claims, msg
}


func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string){
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{"token", signedToken})
	updateObj = append(updateObj, bson.E{"refreshtoken", signedRefreshToken})

	Updated_at, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		log.Panic(err)
		return
	}
	updateObj = append(updateObj, bson.E{"updatedat", Updated_at})
	upsert := true
	filter := bson.M{"userid":userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err = userTokenCollection.UpdateMany(
		ctx,
		filter,
		bson.D{
			{"$set", updateObj},
		},
		&opt,
	)

	defer cancel()

	if err!=nil{
		log.Fatal(err)
		return
	}
	return 
}

