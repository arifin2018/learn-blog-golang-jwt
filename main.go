package main

import (
	"Blog/config"
	"Blog/routes"
)

func main()  {
	db := config.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter(db)

	r.Run() // listen and serve on 0.0.0.0:8080
}