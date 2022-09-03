package controllers

import (
	"github.com/a1nn1997/go-auth/database"
	"github.com/a1nn1997/go-auth/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {  
		return err
	}  //show error during parsing

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 16)   //hash of password

	user := models.User{
		First_name:     data["first_name"],
		Last_name:     	data["last_name"],
		Email:    		data["email"],
		Phone:     		data["phone"],
		Password: password,
		Is_active: 		true,	
	}   //it will stored  user date for model
	user.Created_at,_= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at,_= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))	
	database.DB.Create(&user)	//data stored in db

	return c.JSON(user)		//send json
}

func EditUser(c *fiber.Ctx) error {
	
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
	if(data["password"]!=""){
		password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 16)   //hash of password
		user.Password=password
	}
	
	user1 := models.User{
		First_name:     data["first_name"],
		Last_name:     	data["last_name"],
		Email:    		data["email"],
		Phone:     		data["phone"],
	}
	if(user1.First_name!=""){
		user.First_name=user1.First_name
	}
	if(user1.Last_name!=""){
		user.Last_name=user1.Last_name
	}
	if(user1.Email!=""){
		user.Email=user1.Email
	}
	if(user1.Phone!=""){
		user.Phone=user1.Phone
	}
	user.Updated_at,_= time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))	
	database.DB.Save(&user)
	return c.JSON(user)		//send json
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)  //search by email first value

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}//if user not found

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}  //incorrect password

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	}) //creaate jwt tken with 1day expire time

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}	// token not found in cookies

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}  // cookies updation

	c.Cookie(&cookie) //update cookies

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func User(c *fiber.Ctx) error {
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

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	// clear cookies 
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}