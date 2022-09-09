package models

import(
	"time"
)

type Comment struct {
	Id       		uint   				`json:"id"`
	Text			string				`json:"text" validate:"required min=1"`
	Created_at		time.Time			`json:"created_at"`
	Updated_at		time.Time			`json:"updated_at"`
	User_Id			string				`json:"username" validate:"required"`
	Post_id			uint				`json:"postId"  validate:"required"`
}