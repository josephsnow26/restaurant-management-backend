// @title Restaurant Management API
// @version 1.0
// @description This is a restaurant management server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
package main

import (
	"golang_restaurant_management/controllers"
	"golang_restaurant_management/database"
	_ "golang_restaurant_management/docs"
	"golang_restaurant_management/routes"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	// Mongo collections
	usersCollection := database.OpenCollection(database.Client, "users")
	controllers.SetUsersCollection(usersCollection)

	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Routes
	routes.UserRoutes(router)

	router.Run(":" + port)
}
