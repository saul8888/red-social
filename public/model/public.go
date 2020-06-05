package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetLimit struct {
	Limit  int `json:"limit" form:"limit" query:"limit"`
	Offset int `json:"offset" form:"offset" query:"offset"`
}

type Public struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID    primitive.ObjectID `bson:"userId" json:"userId,omitempty"`
	Message   string             `bson:"message" json:"message,omitempty"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
	//Coordinates CoordinatesLo      `bson:"coordinates" json:"coordinates"`
}

type PublicUpdate struct {
	//UserID  string    `bson:"userId" json:"userId,omitempty"`
	Message string `bson:"message" json:"message,omitempty"`
}

var Publicupdate = map[string]interface{}{
	"userid":    "",
	"message":   "",
	"updatedat": time.Now(),
}
