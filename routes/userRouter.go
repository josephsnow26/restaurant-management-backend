package routes

import (
	"github.com/gin-gonic/gin"
	controller "golang_restaurant_management/controllers"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/users", controller.GetUsers)
	router.POST("/users", controller.CreateUser)

}
