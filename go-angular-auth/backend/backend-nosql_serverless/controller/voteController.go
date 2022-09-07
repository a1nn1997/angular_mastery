package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/a1nn1997/redditclone/database"
	"github.com/a1nn1997/redditclone/models"
	utils "github.com/a1nn1997/redditclone/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)
var voteCollection *mongo.Collection = database.OpenCollection(database.Client, "vote")

func VotePost() gin.HandlerFunc {
	return func(c *gin.Context) {
				getToken := c.Request.Header["Token"][0]
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"token": getToken}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"type":-1,"error": err.Error()})
			return
		}
		if err := utils.AllowedType(c, "ADMIN", "USER"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"type":0,"error": err.Error()})
			return
		}
		var vote models.Vote
		if err :=c.BindJSON(&vote); err !=nil{

			c.JSON(http.StatusBadRequest, gin.H{"type":1,"error":err.Error()})
			return
		}
		vote.Id=primitive.NewObjectID()
		vote.User_Id=	user.User_id
		vote.Vote_id=vote.Id.Hex()
		validationErr := validate.Struct(vote)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"type":2,"error":validationErr.Error()})
			return
		}
		vote.Created_at, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"type":3,"error": "error occured while updating time"})
		}
		vote.Updated_at, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"type":4,"error": "error occured while updating time"})
		}		
		_, insertErr := voteCollection.InsertOne(ctx, vote)
		if insertErr != nil{
			msg :=fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"type":5,"status":msg,"error":insertErr})
			return
		}
		var post models.Post
		err = postCollection.FindOne(ctx, bson.M{"post_id":vote.Post_id}).Decode(&post)
		defer cancel()
		if err !=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"type":vote.Post_id,"error":err.Error()})
			return
		}
		postvote := post.Vote
		if(vote.Type=="UPVOTE"){
			postvote+=1
		}else{
			postvote-=1
		}		
		edited := bson.M{"vote": postvote}
 
		_, err = postCollection.UpdateOne(ctx, bson.M{"post_id": vote.Post_id}, bson.M{"$set": edited})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		_= postCollection.FindOne(ctx, bson.M{"post_id":vote.Post_id}).Decode(&post)
		type Votestruct struct{
			Vote_Outcome 	models.Vote		`json:"vote_outcome"`
			Post_Outcome	models.Post		`json:"post_outcome"`
		}
		vs := Votestruct{Vote_Outcome:vote,Post_Outcome:post}
		c.JSON(http.StatusOK, vs)
	}
}