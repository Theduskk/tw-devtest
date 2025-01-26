package main

import (
	apis "tw-devtask/apis"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin routers for all endpoints
	// Runs on default port 3000
	router := gin.Default()
	router.POST("/subscribe", apis.HandlerSubscribe)
	router.POST("/unsubscribe", apis.HandlerUnsubscribe)
	router.GET("/getSubscribedAccounts", apis.HandlerGetSubscribedAccounts)
	router.GET("/getCurrentBlock", apis.HandlerGetCurrentBlock)
	router.GET("/getTransactions", apis.HandlerGetTransactions)
	router.GET("/getAllTransactions", apis.HandlerGetAllTransactions)
	router.Run("0.0.0.0:3000")
}
