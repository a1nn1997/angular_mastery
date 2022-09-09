package models

import(
	"time"
)

type SubReddit struct {
	Id       		uint   				`json:"id"`
	Title			string				`json:"title" validate:"required,min=1"`
	Description		string				`json:"description" validate:"required,min=1"`
	NumberOfPost	uint32				`json:"numberOfPosts" validate:"required"`
	Created_at		time.Time			`json:"created_at"`
	Updated_at		time.Time			`json:"updated_at"`
	User_Id			string              `json:"user_id" validate:"required"`								
}