package config

import (
	"Blog/helpers"
	"Blog/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var username = helpers.Getenv("username", "root")
var password = helpers.Getenv("password", "123456")
var host = helpers.Getenv("host", "127.0.0.1")
var port = helpers.Getenv("port", "3306")
var database = helpers.Getenv("database", "blogGO")

func ConnectDB() *gorm.DB {
	// username := "root"
	// password := "123456"
	// host 	 := "127.0.0.1"
	// port 	 := "3306"
	// database := "blogGO"

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",username,password,host,port,database)
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
		// panic("err.Error()")
	}

	db.AutoMigrate(&models.User{},&models.Post{}, &models.Tag{},&models.Image{},&models.Comment{})

	if err := db.Migrator().DropColumn(&models.Comment{}, "image_url"); err != nil{
		// Do whatever you want to do!
		log.Print("ERROR: We expect the description column to be drop-able")
	}
	return db
}