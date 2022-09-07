package main

import (
	"log"
	"os"

	routes "github.com/a1nn1997/redditclone/routes"
	"github.com/a1nn1997/redditclone/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	dynaClient dynamodbiface.DynamoDBAPI
)   //var for apis

func main() {
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal("error in loading env file")
	}
	port :=os.Getenv("PORT")
	if port == "" {
		port="8081"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	
	router.GET("/api-1", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"Acess granted for api-1"})
	})

	router.GET("/api-2", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"Acess granted for api-2"})
	})
	router.Run(":"+port)

	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(&aws.Config{ //region where actual session is created 
		Region: aws.String(region)})

	if err != nil {
		return
	}

	dynaClient = dynamodb.New(awsSession)  //create dynamodb clientwith aws session
	lambda.Start(handler)
}

const tableName = "vote_serverless"  //create tables 

//handler for apigateway
func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "POST":
		return handlers.CreateVote(req, tableName, dynaClient)
		default:
		return handlers.UnhandledMethod()  //url like api
	}
}
