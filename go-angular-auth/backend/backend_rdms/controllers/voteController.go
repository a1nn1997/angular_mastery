package controllers

import (
	"github.com/a1nn1997/go-auth/database"
	"github.com/a1nn1997/go-auth/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"time"
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
	postid,err:=strconv.Atoi(data["post_id"])
	if err != nil {
		return c.Status(400).JSON("unable to parse post id")
	}
	vote := models.Vote{
		Type: data["vote_type"],
		}
	vote.User_Id=user.Id
	vote.Post_id=uint(postid)
	
	var post models.Post
		err = findPost(postid, &post)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	
	if(data["vote_type"]=="UPVOTE"){
	post.Vote=post.Vote+1
	}else{
		post.Vote=post.Vote-1	
	}
	database.DB.Save(&post)
	vote.Created_at,_= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	vote.Updated_at,_= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	database.DB.Create(&vote)
	type Votestruct struct{
		Vote_Outcome 	models.Vote		`json:"vote_outcome"`
		Post_Outcome	models.Post		`json:"post_outcome"`
	}
	vs	:= Votestruct{Vote_Outcome:vote,Post_Outcome:post}
	return c.JSON(vs)

}
