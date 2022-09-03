package models

import(
	"time"
)

type Vote struct {
	Id       		uint   				`json:"id"`
	Type			string				`json:"vote_type" validate:"required"`
	Created_at		time.Time			`json:"created_at"`
	Updated_at		time.Time			`json:"updated_at"`
	User_Id			uint				`json:"user_id" validate:"required"`				
	Post_id			uint				`json:"post_id" validate:"required"`				
}