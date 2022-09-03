package controllers

import (
	"github.com/a1nn1997/go-auth/database"
	"github.com/a1nn1997/go-auth/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"errors"
	"time"
	"strconv"
)
func CreateResponseComment(comment models.Comment) models.Comment {
	return models.Comment{Id: comment.Id, Text: comment.Text, Created_at: comment.Created_at, Updated_at: comment.Updated_at, User_Id: comment.User_Id , Post_id: comment.Post_id}
}

func findComment(id int, comment *models.Comment) error {
	database.DB.Find(&comment, "id = ?", id)
	if comment.Id == 0 {
		return errors.New("Comment does not exist")
	}
	return nil
}

func AddComment(c *fiber.Ctx) error {
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
	postid,err:=strconv.Atoi(data["post_id"])
	if err != nil {
		return c.Status(400).JSON("unable to parse Comment id")
	}
	comment := models.Comment{
		Text:     data["text"],
	}   //it will stored  user date for model
	comment.Post_id=uint(postid)
	comment.User_Id=user.Id
	comment.Created_at,_= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	comment.Updated_at,_= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))	
	database.DB.Create(&comment)	//data stored in db
	return c.JSON(comment)		//send json
}
func GetComment(c *fiber.Ctx) error {
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
	comments := []models.Comment{}
	database.DB.Find(&comments)
	responseComments := []models.Comment{}
	for _, comment := range comments {
		responseComment := CreateResponseComment(comment)
		responseComments = append(responseComments, responseComment)
	}

	return c.Status(200).JSON(responseComments)
}
func EditComment(c *fiber.Ctx) error {
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

	var comment models.Comment

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findComment(id, &comment)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := c.BodyParser(&data); err != nil {  
		return err
	} 
   //it will stored  user date for model
		comment.Text=data["text"]	
	comment.Updated_at,_= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	return c.JSON(comment)
}
func GetCommentByCommentId(c *fiber.Ctx) error {
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

	var comment models.Comment

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findComment(id, &comment)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseComment := CreateResponseComment(comment)

	return c.Status(200).JSON(responseComment)	
}
func GetCommentByUserId(c *fiber.Ctx) error {
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
	comments:= []models.Comment{}
	database.DB.Where("user_id = ?", user.Id).Find(&comments)
	responseComments := []models.Comment{}
	for _, comment := range comments {
		responseComment := CreateResponseComment(comment)
		responseComments = append(responseComments, responseComment)
	}

	return c.Status(200).JSON(responseComments)
}

func GetCommentByPostId(c *fiber.Ctx) error {
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
	comments:= []models.Comment{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}
	database.DB.Where("post_id = ?", id).Find(&comments)
	responseComments := []models.Comment{}
	for _, comment := range comments {
		responseComment := CreateResponseComment(comment)
		responseComments = append(responseComments, responseComment)
	}

	return c.Status(200).JSON(responseComments)
}

func GetCommentById(c *fiber.Ctx) error {
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
	comments:= []models.Comment{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}
	database.DB.Where("id = ?", id).Find(&comments)
	responseComments := []models.Comment{}
	for _, comment := range comments {
		responseComment := CreateResponseComment(comment)
		responseComments = append(responseComments, responseComment)
	}

	return c.Status(200).JSON(responseComments)
}

