package controllers

import (
	"errors"
	"strings"
	"time"
	"github.com/a1nn1997/go-auth/database"
	"github.com/a1nn1997/go-auth/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func CreateResponsePost(post models.Post) models.Post {
	return models.Post{Id: post.Id, Description: post.Description, Post_Name: post.Post_Name,Url: post.Url,Vote: post.Vote, Created_at: post.Created_at, Updated_at: post.Updated_at, User_Id: post.User_Id}
}

func CreatePost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")  //store cookies from jwt

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})  //cookies matching

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}  //cookies matching

	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User  //user
	database.DB.Where("id = ?", claims.Issuer).First(&user) //get user by id 
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {  
		return err
	}  //show error during parsing
	/*var subreddit models.SubReddit
	database.DB.Where(&subreddit, "title = ?", data["subreddit_id"]).First(&subreddit)
	subredditid:=subreddit.Id,data["subreddit_id"]
	if err != nil {
		return c.Status(400).JSON("unable to parse post id")
	}*/
	post := models.Post{
		Description:     data["description"],
		Post_Name:     	data["title"],
		Url:    		data["url"],
		Vote:			0,
		Upvote:			false,
		Downvote:		false,
		Comment_count:	0,

	}   //it will stored  user date for model
	post.Subreddit_id=data["subredditName"]
	post.User_Id=strings.Title(user.First_name)+" "+strings.Title(user.Last_name)
	post.Created_at,_= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	post.Updated_at,_= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))	
	database.DB.Create(&post)	//data stored in db
	var subreddit models.SubReddit
	database.DB.Where("title = ?", data["subredditName"]).First(&subreddit)
	subreddit.NumberOfPost=subreddit.NumberOfPost+1
	database.DB.Save(&subreddit)


	return c.JSON(post)		//send json
}

func findPost(id int, post *models.Post) error {
	database.DB.Find(&post, "id = ?", id)
	if post.Id == 0 {
		return errors.New("post does not exist")
	}
	return nil
}

func EditPost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")  //store cookies from jwt

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})  //cookies matching

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}  //cookies matching

	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User  //user
	database.DB.Where("id = ?", claims.Issuer).First(&user) //get user by id 
	var data map[string]string
	
	id, err := c.ParamsInt("id")

	var post models.Post

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findPost(id, &post)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := c.BodyParser(&data); err != nil {  
		return err
	} 
	post1 := models.Post{
		Description:     data["description"],
		Post_Name:     	data["title"],
		Url:    		data["url"],
	}   //it will stored  user date for model
	post1.Updated_at,_= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if(post1.Description!=""){
		post.Description=post1.Description
	}
	if(post1.Post_Name!=""){
		post.Post_Name=post1.Post_Name
	}
	if(post1.Url!=""){
		post.Url=post1.Url	}
	post.Updated_at,_= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	return c.JSON(post)
}

func GetPost(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")  //store cookies from jwt

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})  //cookies matching

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}  //cookies matching

	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User  //user
	database.DB.Where("id = ?", claims.Issuer).First(&user) //get user by id 
	posts := []models.Post{}
	database.DB.Find(&posts)
	responsePosts := []models.Post{}
	for _, post := range posts {
		responsePost := CreateResponsePost(post)
		responsePosts = append(responsePosts, responsePost)
	}

	return c.Status(200).JSON(responsePosts)
}
func GetPostById(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")  //store cookies from jwt

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})  //cookies matching

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}  //cookies matching

	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User  //user
	database.DB.Where("id = ?", claims.Issuer).First(&user) //get user by id 
	id, err := c.ParamsInt("id")

	var post models.Post

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findPost(id, &post)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responsePost := CreateResponsePost(post)

	return c.Status(200).JSON(responsePost)	
}

func PostByUserName(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")  //store cookies from jwt

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})  //cookies matching

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}  //cookies matching

	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User  //user
	database.DB.Where("id = ?", claims.Issuer).First(&user) //get user by id 
	posts:= []models.Post{}
	database.DB.Where("user_id = ?", user.Id).Find(&posts)
	responsePosts := []models.Post{}
	for _, post := range posts {
		responsePost := CreateResponsePost(post)
		responsePosts = append(responsePosts, responsePost)
	}

	return c.Status(200).JSON(responsePosts)
}

func PostBySubRedditId(c *fiber.Ctx) error {
		cookie := c.Cookies("jwt")  //store cookies from jwt

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})  //cookies matching

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}  //cookies matching

	claims := token.Claims.(*jwt.StandardClaims)
	var user models.User  //user
	database.DB.Where("id = ?", claims.Issuer).First(&user) //get user by id 
	posts:= []models.Post{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}
	database.DB.Where("subreddit_id = ?", id).Find(&posts)
	responsePosts := []models.Post{}
	for _, post := range posts {
		responsePost := CreateResponsePost(post)
		responsePosts = append(responsePosts, responsePost)
	}

	return c.Status(200).JSON(responsePosts)
}


