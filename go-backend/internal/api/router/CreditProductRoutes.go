package router

import (
	"app/internal/api/controller"
	"github.com/gin-gonic/gin"
)

func (router *Router) CreditProductRoutes(group *gin.RouterGroup) {
	group.GET("/products", controller.GetProducts)
	group.GET("/products/:productId", controller.GetProductByID)
	group.POST("/customer-leads", controller.CreateCustomerLead)
	group.GET("/customer-leads/:customerLeadId", controller.GetCustomerLeadByID)
	group.DELETE("/customer-leads/:customerLeadId", controller.DeleteCustomerLead)
	group.POST("/product-offers", controller.CreateProductOffer)
	group.GET("/product-offers", controller.GetProductOffers)
	group.GET("/product-offers/:offerId", controller.GetProductOfferByID)
	group.DELETE("/product-offers/:offerId", controller.DeleteProductOffer)
	group.POST("/product-offer-consents", controller.CreateProductOfferConsent)
	group.GET("/product-offer-consents/:consentId", controller.GetProductOfferConsentByID)
	group.DELETE("/product-offer-consents/:consentId", controller.DeleteProductOfferConsent)
	group.POST("/product-application", controller.CreateProductApplication)
	group.GET("/product-application", controller.GetProductApplications)
	group.GET("/product-application/:productApplicationId", controller.GetProductApplicationByID)
	group.DELETE("/product-application/:productApplicationId", controller.DeleteProductApplication)
}
