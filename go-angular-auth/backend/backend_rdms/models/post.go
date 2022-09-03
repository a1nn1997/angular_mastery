package models

import(
	"time"
)	

type Post struct {
	Id       		uint   				`json:"id"`
	Description		string				`json:"description" validate:"required,min=10"`
	Post_Name		string				`json:"title" validate:"required,min=3"`
	Url				string				`json:"url"`
	Created_at		time.Time			`json:"created_at"`
	Updated_at		time.Time			`json:"updated_at"`
	Vote			int					`json:"vote" validate:"required"`
	User_Id			uint				`json:"user_id" validate:"required"`
	Subreddit_id	uint				`json:"subreddit_id"`			
}