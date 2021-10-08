package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id        primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	UserId    string              `bson:"userId" json:"userId"`
	Caption   string              `bson:"caption" json:"caption"`
	ImageUrl  string              `bson:"imageUrl" json:"imageUrl"`
	Timestamp primitive.Timestamp `bson:"timestamp" json:"timestamp"`
}
