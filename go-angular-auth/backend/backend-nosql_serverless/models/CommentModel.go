package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

)

type Comment struct {
	ID                 	primitive.ObjectID   `bson:"_id"`
	Text				string				`json:"text" validate:"required"`
	Created_at			time.Time			`json:"created_at"`
	Updated_at			time.Time			`json:"updated_at"`
	User_id				string				`json:"user_id"`
	Post_id				string				`json:"post_id"`
	Comment_id			string				`json:"comment_id"`
}