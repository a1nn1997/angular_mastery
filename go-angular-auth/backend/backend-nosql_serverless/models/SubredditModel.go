package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"

)	

type Subreddit struct {
	Id       		primitive.ObjectID  	 	`bson:"_id"`		
	Title			string						`json:"title" validate:"required,min=1"`
	Description		string						`json:"description" validate:"required,min=1"`
	Created_at		time.Time					`json:"created_at"`
	Updated_at		time.Time					`json:"updated_at"`
	User_Id			string						`json:"user_id" validate:"required"`
	Subreddit_id	string						`json:"subreddit_id"`			
}