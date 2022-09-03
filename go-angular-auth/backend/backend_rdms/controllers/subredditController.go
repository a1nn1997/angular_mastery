package controllers

import (
	"github.com/a1nn1997/go-auth/database"
	"github.com/a1nn1997/go-auth/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"time"
	"errors"
)

func CreateResponseSubReddit(subreddit models.SubReddit) models.SubReddit {
	return models.SubReddit{Id: subreddit.Id,Description: subreddit.Description, Title: subreddit.Title, Created_at: subreddit.Created_at, Updated_at: subreddit.Updated_at, User_Id: subreddit.User_Id,}
}
func CreateSubreddit(c *fiber.Ctx) error {
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
	subreddit := models.SubReddit{
		Title: data["title"],
		Description: data["description"],
		}
	subreddit.User_Id=user.Id
	subreddit.Created_at,_= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	subreddit.Updated_at,_= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	database.DB.Create(&subreddit)
	return c.JSON(subreddit)

}

func GetSubreddit(c *fiber.Ctx) error {
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
	subreddits := []models.SubReddit{}
	database.DB.Find(&subreddits)
	response_subreddits := []models.SubReddit{}
	for _, subreddit := range subreddits {
		response_subreddit := CreateResponseSubReddit(subreddit)
		response_subreddits = append(response_subreddits, response_subreddit)
	}

	return c.Status(200).JSON(response_subreddits)
}

func findSubReddit(id int, post *models.SubReddit) error {
	database.DB.Find(&post, "id = ?", id)
	if post.Id == 0 {
		return errors.New("post does not exist")
	}
	return nil
}

func GetSubredditById(c *fiber.Ctx) error {
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

	var subreddit models.SubReddit

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findSubReddit(id, &subreddit)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseSubReddit := CreateResponseSubReddit(subreddit)

	return c.Status(200).JSON(responseSubReddit)	
}
