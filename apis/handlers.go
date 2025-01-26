package apis

import (
	"fmt"
	"net/http"

	gin "github.com/gin-gonic/gin"
)

// Handler middlewares to pipe required data from context into smaller functions
func HandlerGetCurrentBlock(c *gin.Context) {
	address := extractAddressFromPayload(c)
	fmt.Printf("address: %s", address)
	if address == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "address not provided"})
		return
	}
	body, err := GetCurrentBlock(address)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("HandlerGetCurrentBlock: body: %s", body)
	c.IndentedJSON(http.StatusOK, body)
}

func HandlerSubscribe(c *gin.Context) {
	address := extractAddressFromPayload(c)
	fmt.Printf("address: %s", address)
	if address == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "address not provided"})
		return
	}
	fmt.Printf("address: %s", address)
	err := Subscribe(address)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("address: %s", address)
	c.IndentedJSON(http.StatusOK, address)
}

func HandlerUnsubscribe(c *gin.Context) {
	address := extractAddressFromPayload(c)
	fmt.Printf("address: %s", address)
	if address == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "address not provided"})
		return
	}
	err := Unsubscribe(address)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func HandlerGetSubscribedAccounts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"accounts": GetSubscribedAccounts()})
}

func HandlerGetTransactions(c *gin.Context) {
	address := extractAddressFromPayload(c)
	if address == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "address not provided"})
		return
	}
	fmt.Printf("address: %s", address)
	transactions, err := GetTransactions(address)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("transactions: %v", transactions)
	c.IndentedJSON(http.StatusOK, gin.H{"transactions": transactions})
}

func HandlerGetAllTransactions(c *gin.Context) {
	transactions, err := GetAllTransactions()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("transactions: %v", transactions)
	c.IndentedJSON(http.StatusOK, gin.H{"transactions": transactions})
}

func extractAddressFromPayload(c *gin.Context) string {
	if address, ok := c.GetQuery("address"); ok {
		return address
	}
	return ""
}
