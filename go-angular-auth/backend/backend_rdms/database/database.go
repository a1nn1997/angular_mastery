package database

import (
	"github.com/a1nn1997/go-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
func Connect() {
	dsn := "root:123@tcp(127.0.0.1:3500)/go_ang_auth?charset=utf8mb4&parseTime=True&loc=Local"    //db setup  mysql  tcp is port of mysql
    connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB = connection
	connection.AutoMigrate(&models.User{}, &models.Comment{}, &models.Post{}, &models.SubReddit{}, &models.Vote{})
}