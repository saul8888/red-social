package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetLimit struct {
	Limit  int `json:"limit" form:"limit" query:"limit"`
	Offset int `json:"offset" form:"offset" query:"offset"`
}

type Follow struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID      primitive.ObjectID `bson:"userId" json:"userId,omitempty"`
	FollowingID primitive.ObjectID `bson:"followingId" json:"followingId,omitempty"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}
