package routes

import (
	"github.com/gin-gonic/gin"
	"golang_restaurant_management/controllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoice", controllers.GetInvoices())
	incomingRoutes.GET("/invoice/:invoice_id", controllers.GetInvoice())
	incomingRoutes.POST("/invoice", controllers.CreateInvoice())
	incomingRoutes.PATCH("invoice/:invoice_id", controllers.UpdateInvoice())

}
