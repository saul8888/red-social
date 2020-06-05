package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetLimit struct {
	Limit  int `json:"limit" form:"limit" query:"limit"`
	Offset int `json:"offset" form:"offset" query:"offset"`
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserName  string             `bson:"userName" json:"userName" validate:"required"`
	FirstName string             `bson:"firstName" json:"firstName" validate:"required"`
	LastName  string             `bson:"lastName" json:"lastName" validate:"required"`
	Email     string             `bson:"email" json:"email" validate:"required,email"`
	Password  string             `bson:"password,omitempty" json:"password" validate:"required"`
	Avatar    string             `bson:"avatar,omitempty" json:"avatar"`
	Front     string             `bson:"front,omitempty" json:"front"`
	Biografia string             `bson:"biografia,omitempty" json:"biografia"`
	Location  string             `bson:"location" json:"location" validate:"required"`
	Website   string             `bson:"website,omitempty" json:"website"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
	//Birthday  time.Time          `bson:"birthday" json:"birthday"`
}

type UserUpdate struct {
	UserName  string `bson:"userName" json:"userName" validate:"required"`
	FirstName string `bson:"firstName" json:"firstName" validate:"required"`
	LastName  string `bson:"lastName" json:"lastName" validate:"required"`
	Email     string `bson:"email" json:"email" validate:"required,email"`
	Password  string `bson:"password,omitempty" json:"password" validate:"required"`
	Avatar    string `bson:"avatar,omitempty" json:"avatar"`
	Front     string `bson:"front,omitempty" json:"front"`
	Biografia string `bson:"biografia" json:"biografia"`
	Location  string `bson:"location" json:"location" validate:"required"`
	Website   string `bson:"website" json:"website"`
}

var Userupdate = map[string]interface{}{
	"userName":  "",
	"firstName": "",
	"lastName":  "",
	"email":     "",
	"password":  "",
	"Avatar":    "",
	"front":     "",
	"biografia": "",
	"location":  "",
	"website":   "",
	"updatedat": time.Now(),
}
