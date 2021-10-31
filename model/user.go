package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `form:"name" json:"name"`
	Email string `form:"email" json:"email," binding:"required"`
	Password string `form:"password" json:"password,omitempty" binding:"required"`
	CreatedAt time.Time `form:"createdAt" json:"createdAt,omitempty"`
}
