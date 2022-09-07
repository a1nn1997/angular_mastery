package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vote struct {
	Id     			primitive.ObjectID 		`bson:"id"`
	Type			string					`json:"vote_type" validate:"required,eq=UPVOTE|eq=DOWNVOTE"`
	Created_at		time.Time				`json:"created_at"`
	Updated_at		time.Time				`json:"updated_at"`
	User_Id			string					`json:"user_id" validate:"required"`				
	Post_id			string					`json:"post_id" validate:"required"`				
	Vote_id			string					`json:"vote_id"`
}
