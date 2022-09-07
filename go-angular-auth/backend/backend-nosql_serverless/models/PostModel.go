package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"

)	

type Post struct {
	Id       		primitive.ObjectID  	 	`bson:"_id"`		
	Description		string						`json:"description" validate:"required,min=10"`
	Post_Name		string						`json:"title" validate:"required,min=3"`
	Url				string						`json:"url"`
	Created_at		time.Time					`json:"created_at"`
	Updated_at		time.Time					`json:"updated_at"`
	Vote			int							`json:"vote"`
	User_Id			string						`json:"user_id" validate:"required"`
	Post_Id			string						`json:"post_id" validate:"required"`	
	Subreddit_id	string						`json:"subreddit_id"`			
}