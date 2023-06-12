package main

import (
	"Blog/config"
)

func main()  {
	db := config.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
}