package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	TodoContent string             `form:"todo" json:"todo" binding:"required"`
	CreatedAt   time.Time          `form:"createdAt" json:"createdAt,omitempty"`
	UserId 			primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty"`
}
