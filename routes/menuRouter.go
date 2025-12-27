package routes

import (
	"github.com/gin-gonic/gin"
	"golang_restaurant_management/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menu", controllers.GetMenus())
	incomingRoutes.GET("/menu/:menu_id", controllers.GetMenu())
	incomingRoutes.POST("/menu", controllers.CreateMenu())
	incomingRoutes.PATCH("menu/:menu_id", controllers.UpdateMenu())

}
