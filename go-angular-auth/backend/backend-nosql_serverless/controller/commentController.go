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

var commentCollection *mongo.Collection = database.OpenCollection(database.Client, "comment")

func GetComment() gin.HandlerFunc {
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
		result, err := commentCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		var allcomments []bson.M
		if err = result.All(ctx, &allcomments); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, allcomments[0])
	}
}

func GetCommentById() gin.HandlerFunc {
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
		var comment models.Comment
		commentId := c.Param("id")
		err = commentCollection.FindOne(ctx, bson.M{"comment_id":commentId}).Decode(&comment)
		defer cancel()
		if err !=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
		c.JSON(http.StatusOK, comment)
	}
}

func AddComment() gin.HandlerFunc {
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
		var comment models.Comment
		if err :=c.BindJSON(&comment); err !=nil{

			c.JSON(http.StatusBadRequest, gin.H{"type":1,"error":err.Error()})
			return
		}
		comment.ID=primitive.NewObjectID()
		comment.User_id=	user.User_id
		comment.Comment_id=comment.ID.Hex()
		validationErr := validate.Struct(comment)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"type":2,"error":validationErr.Error()})
			return
		}
		comment.Created_at, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"type":3,"error": "error occured while updating time"})
		}
		comment.Updated_at, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"type":4,"error": "error occured while updating time"})
		}		
		_, insertErr := commentCollection.InsertOne(ctx, comment)
		if insertErr != nil{
			msg :=fmt.Sprintf("User item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"type":5,"status":msg,"error":insertErr})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, comment)
	}
}

func EditComment() gin.HandlerFunc {
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

		type Comment1 struct {
			Text			string						`json:"text" validate:"required min=1"`
			Updated_at		time.Time					`json:"updated_at"`
		}
		var comment models.Comment
		var comment1 Comment1
		if err := c.BindJSON(&comment1); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}

		if comment1.Text!=""{		
			edited := bson.M{"Text": comment1.Text}
 
			_, err := commentCollection.UpdateOne(ctx, bson.M{"comment_id": ObjectID}, bson.M{"$set": edited})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err})
				return
			}
		}
	
		comment1.Updated_at, err = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while updating time"})
		}
		edited := bson.M{"updated_at": comment1.Updated_at}
		_, err = commentCollection.UpdateOne(ctx, bson.M{"comment_id": ObjectID}, bson.M{"$set": edited})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err})
			return
		}
		defer cancel()
		_ = commentCollection.FindOne(ctx, bson.M{"comment_id":ObjectID}).Decode(&comment)		
		c.JSON(http.StatusOK, comment)
	}
}

func GetCommentByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		getToken:=c.Request.Header["Token"][0]
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var comment []*models.Comment
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
		cur,err := commentCollection.Find(context.TODO(), bson.D{{"user_id",user.User_id}}, findOptions)
		if err !=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
		err = cur.All(ctx, &comment)
		defer cancel()
		
		c.JSON(http.StatusOK,  comment)
	}
}

func GetCommentByPostId() gin.HandlerFunc {
	return func(c *gin.Context) {
		getToken:=c.Request.Header["Token"][0]
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		id :=c.Param("id")
		var comment []*models.Comment
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
		cur,err := commentCollection.Find(context.TODO(), bson.D{{"post_id",id}}, findOptions)
		if err !=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
			return
		}
		err = cur.All(ctx, &comment)
		defer cancel()
		
		c.JSON(http.StatusOK,  comment)
	}
}

