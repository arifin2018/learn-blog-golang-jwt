package main

import (
	"Blog/config"
	"Blog/docs"
	"Blog/routes"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main()  {
	//programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API by Nur Arifin"
	docs.SwaggerInfo.Description = "This is a sample server Blog."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	db := config.ConnectDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	r := routes.SetupRouter(db)

	r.Run() // listen and serve on 0.0.0.0:8080
}