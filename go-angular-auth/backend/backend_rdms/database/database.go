package database

import (
	"github.com/a1nn1997/go-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *gorm.DB

// DBConfig represents db configuration
func Connect() {
	err := godotenv.Load(".env")
	if err!=nil{
		log.Fatal("error loading .env file")
	}
	dsn := os.Getenv("DATABASE_URL")    //db setup  mysql  tcp is port of mysql
    connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = connection
	connection.AutoMigrate(&models.User{}, &models.Comment{}, &models.Post{}, &models.SubReddit{}, &models.Vote{})
}