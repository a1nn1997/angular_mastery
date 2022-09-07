package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	//"reflect"
	"strconv"
	//"strings"
	"time"

	"github.com/a1nn1997/redditclone/database"
	"github.com/a1nn1997/redditclone/models"
	utils "github.com/a1nn1997/redditclone/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var subredditCollection *mongo.Collection =database.OpenCollection(database.Client, "subreddit")

func GetSubreddit() gin.HandlerFunc{
	return func(c *gin.Context){
		getToken := c.Request.Header["Token"][0]
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"token": getToken}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := utils.AllowedType(c, "ADMIN", "USER"); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}
		page, err1 := strconv.Atoi(c.Query("page"))
		if err1 != nil || page < 1 {
			page = 1
		}
		startIndex := (page - 1) * recordPerPage
		startIndex, err = strconv.Atoi("startIndex")

		matchStage := bson.D{{"$match", bson.D{{}}}}
		groupStage := bson.D{{"$group",
			bson.D{{"_id", bson.D{{"_id", "null"}}},
				{"total_count", bson.D{{"$sum", 1}}},
				{"data", bson.D{{"$push", "$$ROOT"}}}}}}
		projectStage := bson.D{{"$project",
			bson.D{{"_id", 0}, {"total_count", 1},
				{"user_items", bson.D{{"$slice", []interface{}{"$data", startIndex, recordPerPage}}}},
			}}}
		result, err := subredditCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		var allsubreddits []bson.M
		if err = result.All(ctx, &allsubreddits); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allsubreddits[0])	
	}
}

func GetSubredditById() gin.HandlerFunc{
	return func(c *gin.Context){
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
		var subreddit models.Subreddit
		subredditId := c.Param("id")
		err = subredditCollection.FindOne(ctx, bson.M{"subreddit_id":subredditId}).Decode(&subreddit)
		defer cancel()
		if err !=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK, subreddit)
	}
}

func CreateSubreddit() gin.HandlerFunc{
	return func(c *gin.Context){
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
		var subreddit models.Subreddit
		if err :=c.BindJSON(&subreddit); err !=nil{

			c.JSON(http.StatusBadRequest, gin.H{"type":1,"error":err.Error()})
			return
		}
		subreddit.Id=primitive.NewObjectID()
		subreddit.User_Id=	user.User_id
		subreddit.Subreddit_id=subreddit.Id.Hex()
		validationErr := validate.Struct(subreddit)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"type":2,"error":validationErr.Error()})
			return
		}
		subreddit.Created_at, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"type":3,"error": "error occured while updating time"})
		}
		subreddit.Updated_at, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"type":4,"error": "error occured while updating time"})
		}		
		_, insertErr := subredditCollection.InsertOne(ctx, subreddit)
		if insertErr != nil{
			msg :=fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"type":5,"status":msg,"error":insertErr})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, subreddit)
	}
}

