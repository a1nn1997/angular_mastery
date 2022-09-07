package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/a1nn1997/redditclone/database"
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct{
	Email 		string
	First_name 	string
	Last_name 	string
	Uid 		string
	User_type	string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")  //search collection

var SECRET_KEY string = os.Getenv("SECRET_KEY")   //secret key not aval right now

func GenenrateAllTokens(email string, firstName string, lastName string, userType string, uid string) (signedToken string, signedRefreshToken string, err error ){
	claims := &SignedDetails{
		Email : email,
		First_name : firstName,
		Last_name : lastName,
		Uid: uid,
		User_type: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},  //set expire date of token
	}
	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	token,_ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY)) //set token
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if(err != nil){
		log.Panic(err)
		return
	}
	return token, refreshToken, err
}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string){
	ctx, cancel :=context.WithTimeout(context.Background(),100*time.Second)
	var updateObj primitive.D
	updateObj = append(updateObj, bson.E{"token", signedToken})
	updateObj = append(updateObj, bson.E{"refresh_token",signedRefreshToken})

	Updated_at,err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		log.Panic(err)
			return
	}
	updateObj= append(updateObj,bson.E{"updated_at",Updated_at})  //append the updated at object
	upsert := true
	filter :=bson.M{"user_id":userId}	//find user id
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err = userCollection.UpdateOne( ctx, filter,    //update object in user collection
		bson.D{	{"$set", updateObj}, }, &opt,)
		defer cancel()
	if err != nil {
		log.Panic(err)
		return
	}
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string){
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token)(interface{}, error){
				return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		msg=err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok{
		msg = fmt.Sprintf("token is invalid")
		msg  = err.Error()
		return
	}
	if claims.ExpiresAt <time.Now().Local().Unix(){
		msg = fmt.Sprintf("token is expired")
		msg=err.Error()
		return
	}
	return claims, msg
}
