package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson: "_id"`
	Username     *string            `json: "username" validate:"required,min=2,max=100"`
	Password     *string  		    `json: "password" validate:"required,min=8,max=100,regexp=^(?=.*[A-Z])(?=.*[!@#$&*])(?=.*[0-9]).*$"`
	UserID       string 		    `json: "userid"`
	UserType     string             `json: "usertype" validate:"required,default=USER,eq=ADMIN|eq=USER"`
	OrgID        string             `json: "orgid"`
	CreatedAt    time.Time		    `json: "createdat"`
	UpdatedAt 	 time.Time		    `json: "updatedat"`
}

type UserTokens struct {
	ID         primitive.ObjectID `bson:"_id"`
	UserID     string            `json:"userid"`
	Token         *string		  `json: "token"`
	RefreshToken *string 	 	  `json: "refreshtoken"`
	CreatedAt   time.Time          `json:"createdat"`
	UpdatedAt   time.Time          `json:"updatedat"`
}


type Organization struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      *string            `json:"name" validate:"required,min=2,max=100"`
	CreatedAt time.Time          `json:"createdat"`
	UpdatedAt time.Time          `json:"updatedat"`
}
