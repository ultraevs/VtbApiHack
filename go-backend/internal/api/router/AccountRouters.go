package router

import (
	"app/internal/api/controller"
	"github.com/gin-gonic/gin"
)

func (router *Router) AccountRoutes(group *gin.RouterGroup) {
	group.POST("/account-consents", controller.CreateAccountAccessConsent)
	group.GET("/account-consents/:consentId", controller.GetAccountAccessConsent)
	group.DELETE("/account-consents/:consentId", controller.DeleteAccountAccessConsent)
	group.GET("/account-consents/:consentId/retrieval-grant", controller.GetAccountAccessRetrievalGrant)
	group.GET("/accounts", controller.GetAccounts)
	group.GET("/accounts/:accountId", controller.GetAccount)
	group.GET("/accounts/:accountId/balance", controller.GetAccountBalance)
	group.GET("/balances", controller.GetBalances)
	group.GET("/accounts/:accountId/transaction", controller.GetAccountTransactions)
	group.GET("/transactions", controller.GetTransactions)
	group.POST("/statements/:accountId", controller.CreateStatement)
	group.GET("/accounts/:accountId/statements/:statementId", controller.GetStatementById)
	group.GET("/statements", controller.GetAllStatements)
}
