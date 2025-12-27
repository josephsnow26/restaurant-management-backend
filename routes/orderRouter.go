package routes

import (
	"github.com/gin-gonic/gin"
	"golang_restaurant_management/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/order", controllers.GetOrders())
	incomingRoutes.GET("/order/:order_id", controllers.GetOrder())
	incomingRoutes.POST("/order", controllers.CreateOrder())
	incomingRoutes.PATCH("order/:order_id", controllers.UpdateOrder())

}
