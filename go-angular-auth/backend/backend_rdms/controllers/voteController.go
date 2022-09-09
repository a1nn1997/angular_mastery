package controllers

import (
	"github.com/a1nn1997/go-auth/database"
	"github.com/a1nn1997/go-auth/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func VotePost(c *fiber.Ctx) error {
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
	postid,err:=strconv.Atoi(data["postId"])
	if err != nil {
		return c.Status(400).JSON("unable to parse post id")
	}
	votetype,err:=strconv.Atoi(data["voteType"])
	if err != nil {
		return c.Status(400).JSON("unable to parse voteType")
	}
	var post models.Post
	err = findPost(postid, &post)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	
	if(votetype==0){
	post.Vote=post.Vote+1
	post.Upvote=true
	post.Downvote=false
	}else{
		post.Vote=post.Vote-1
		post.Upvote=false
		post.Downvote=true	
	}
	database.DB.Save(&post)
	return c.JSON(post)

}
