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
	"go.mongodb.org/mongo-driver/mongo/options"
)

var postCollection *mongo.Collection = database.OpenCollection(database.Client, "post")

func GetPost() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		result, err := postCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		var allposts []bson.M
		if err = result.All(ctx, &allposts); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allposts[0])
	}
}

func GetPostById() gin.HandlerFunc {
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
		var post models.Post
		postId := c.Param("id")
		err = postCollection.FindOne(ctx, bson.M{"post_id":postId}).Decode(&post)
		defer cancel()
		if err !=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK, post)
	}
}

func CreatePost() gin.HandlerFunc {
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
		var post models.Post
		if err :=c.BindJSON(&post); err !=nil{

			c.JSON(http.StatusBadRequest, gin.H{"type":1,"error":err.Error()})
			return
		}
		post.Vote=0
		post.Id=primitive.NewObjectID()
		post.User_Id=	user.User_id
		post.Post_Id=post.Id.Hex()
		validationErr := validate.Struct(post)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"type":2,"error":validationErr.Error()})
			return
		}
		post.Created_at, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"type":3,"error": "error occured while updating time"})
		}
		post.Updated_at, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"type":4,"error": "error occured while updating time"})
		}		
		_, insertErr := postCollection.InsertOne(ctx, post)
		if insertErr != nil{
			msg :=fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"type":5,"status":msg,"error":insertErr})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, post)
	}
}

func EditPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		getToken:=c.Request.Header["Token"][0]
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"token":getToken}).Decode(&user)
		if err !=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
			if err :=utils.AllowedType(c, "ADMIN","USER"); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ObjectID := c.Param("id")

		type Post1 struct {
			Description		string						`json:"description" validate:"required,min=10"`
			Post_Name		string						`json:"title" validate:"required,min=3"`
			Url				string						`json:"url"`
			Updated_at		time.Time					`json:"updated_at"`
		}
		var post models.Post
		var post1 Post1
		if err := c.BindJSON(&post1); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		if post1.Description!=""{		
			edited := bson.M{"description": post1.Description}
 
			_, err := postCollection.UpdateOne(ctx, bson.M{"post_id": ObjectID}, bson.M{"$set": edited})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err})
				return
			}
		}
		
		if post1.Post_Name!=""{		
			edited := bson.M{"post_name": post1.Post_Name}
 
			_, err := postCollection.UpdateOne(ctx, bson.M{"post_id": ObjectID}, bson.M{"$set": edited})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err})
				return
			}
		}
		if post1.Url!=""{		
			edited := bson.M{"url": post1.Url}
 
			_, err := postCollection.UpdateOne(ctx, bson.M{"post_id": ObjectID}, bson.M{"$set": edited})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err})
				return
			}
		}
		post1.Updated_at, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while updating time"})
		}
		edited := bson.M{"updated_at": post1.Updated_at}
		_, err = postCollection.UpdateOne(ctx, bson.M{"post_id": ObjectID}, bson.M{"$set": edited})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		defer cancel()
		_ = postCollection.FindOne(ctx, bson.M{"post_id":ObjectID}).Decode(&post)		
		c.JSON(http.StatusOK, post)		
	}
}

func PostByUserName() gin.HandlerFunc {
	return func(c *gin.Context) {
		getToken:=c.Request.Header["Token"][0]
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var post []*models.Post
		err := userCollection.FindOne(ctx, bson.M{"token":getToken}).Decode(&user)
		if err !=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
			if err :=utils.AllowedType(c, "ADMIN","USER"); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		findOptions := options.Find()
		cur,err := postCollection.Find(context.TODO(), bson.D{{"user_id",user.User_id}}, findOptions)
		if err !=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
		err = cur.All(ctx, &post)
		defer cancel()
		
		c.JSON(http.StatusOK,  post)
	}
}

func PostBySubRedditId() gin.HandlerFunc {
	return func(c *gin.Context) {
		getToken:=c.Request.Header["Token"][0]
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		id :=c.Param("id")
		var post []*models.Post
		err := userCollection.FindOne(ctx, bson.M{"token":getToken}).Decode(&user)
		if err !=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
			if err :=utils.AllowedType(c, "ADMIN","USER"); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		findOptions := options.Find()
		cur,err := postCollection.Find(context.TODO(), bson.D{{"subreddit_id",id}}, findOptions)
		if err !=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
		err = cur.All(ctx, &post)
		defer cancel()
		
		c.JSON(http.StatusOK,  post)
	}
}
