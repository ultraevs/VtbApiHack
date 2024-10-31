package router

import "github.com/gin-gonic/gin"

func (router *Router) VRPRoutes(group *gin.RouterGroup) {
	group.POST("/vrp-consents", controller.CreateVRPConsent)
	group.GET("/vrp-consents/:consentId", controller.GetVRPConsent)
	group.DELETE("/vrp-consents/:consentId", controller.DeleteVRPConsent)
	group.POST("/vrp-consents/:consentId/funds-confirmation", controller.CreateVRPFundsConfirmation)
	group.POST("/vrp-payments", controller.CreateVRPPayment)
	group.GET("/vrp-payments/:VRPId", controller.GetVRPPayment)
	group.GET("/vrp-payments/:VRPId/payment-details", controller.GetVRPPaymentDetails)
}
