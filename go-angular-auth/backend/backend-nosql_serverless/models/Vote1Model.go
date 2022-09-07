package models
// dynamodb db apis 
import (
	"encoding/json"
	"errors"
	"time"
	"context"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/a1nn1997/redditclone/database"
	"github.com/a1nn1997/redditclone/utils"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var postCollection *mongo.Collection = database.OpenCollection(database.Client, "post")

// var to define errors with dynamodb
var (
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
	ErrorFailedToFetchRecord     = "failed to fetch record"
	ErrorInvalidVoteData         = "invalid vote data"
	ErrorInvalidEmail            = "invalid email"
	ErrorCouldNotMarshalItem     = "could not marshal item"
	ErrorCouldNotDeleteItem      = "could not delete item"
	ErrorCouldNotDynamoPutItem   = "could not dynamo put item"
	ErrorVoteAlreadyExists       = "vote.Vote already exists"
	ErrorVoteDoesNotExist        = "vote.Vote does not exist"
)

type Vote1 struct {
	Id     			string 				`json:"id"`
	Email     		string 				`json:"email"`
	Type			string				`json:"vote_type" validate:"required,eq=UPVOTE|eq=DOWNVOTE"`
	Created_at		time.Time			`json:"created_at"`
	Updated_at		time.Time			`json:"updated_at"`
	User_Id			string				`json:"user_id" validate:"required"`				
	Post_id			string				`json:"post_id" validate:"required"`				
}

type Votestruct struct{
	Vote_Outcome 	*Vote1		`json:"vote_outcome"`
	Post_Outcome	*Post		`json:"post_outcome"`
}

func FetchVote(email, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (*Votestruct, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String(tableName),
	}

	result, err := dynaClient.GetItem(input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	item := new(Vote1)
	item.Created_at,_= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	item.Updated_at,_= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	item.Id = uuid.New().String()
	err = dynamodbattribute.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}
	post := new(Post)
	err = postCollection.FindOne(ctx, bson.M{"post_id":item.Post_id}).Decode(&post)
	if err != nil{
		return nil,errors.New(err.Error())
	}
	defer cancel()
	if(item.Type=="UPVOTE"){
		post.Vote=post.Vote+1
	}else{
		post.Vote=post.Vote-1	
	}
	edited := bson.M{"vote": post.Vote}
 
	_, err = postCollection.UpdateOne(ctx, bson.M{"post_id": item.Post_id}, bson.M{"$set": edited})
	if err != nil{
		return nil,errors.New(err.Error())
	}
	
	vs	:= Votestruct{Vote_Outcome:item,Post_Outcome:post}

	return &vs, nil
}


func CreateVote(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*Votestruct,
	error,
) {
	var u Vote1

	if err := json.Unmarshal([]byte(req.Body), &u); err != nil {
		return nil, errors.New(ErrorInvalidVoteData)
	}
	if !utils.IsEmailValid(u.Email) {
		return nil, errors.New(ErrorInvalidEmail)
	}

	currentVote, _ := FetchVote(u.Email, tableName, dynaClient)
	if currentVote != nil && len(currentVote.Vote_Outcome.Email) != 0 {
		return nil, errors.New(ErrorVoteAlreadyExists)
	}

	av, err := dynamodbattribute.MarshalMap(u)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynaClient.PutItem(input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}
	vs	:= Votestruct{Vote_Outcome:&u,Post_Outcome:currentVote.Post_Outcome}
	return &vs, nil
}
