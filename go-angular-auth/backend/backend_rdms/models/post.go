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
	Comment_count	uint				`json:"commentCount" validate:"required"`				
	Vote			uint				`json:"voteCount" validate:"required"`
	Upvote			bool				`json:"upVote" validate:"required"`
	Downvote		bool				`json:"downVote" validate:"required"`
	User_Id			string				`json:"userName" validate:"required"`
	Subreddit_id	string				`json:"subredditName"`			
}