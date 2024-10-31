package router

import (
	"app/internal/api/controller"
	"github.com/gin-gonic/gin"
)

func (router *Router) PaymentsRoutes(group *gin.RouterGroup) {
	group.POST("/user_create", controller.UserCreate)
	group.POST("/login", controller.Login)
	group.POST("/payment-consents", controller.CreatePaymentConsent)
	group.GET("/payment-consents/:consentId", controller.GetPaymentConsent)
	group.POST("/payments", controller.CreatePayment)
	group.GET("/payments/:paymentId", controller.GetPayment)
}
