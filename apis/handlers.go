package apis

import (
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func HandlerGetCurrentBlock(c *gin.Context) {
	address := extractAddressFromPayload(c)
	if address == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "address not provided"})
		return
	}
	body, err := GetCurrentBlock(address)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, body)
}

func HandlerSubscribe(c *gin.Context) {
	address := extractAddressFromPayload(c)
	if address == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "address not provided"})
		return
	}
	err := Subscribe(address)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, address)
}

func HandlerUnsubscribe(c *gin.Context) {
	address := extractAddressFromPayload(c)
	if address == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "address not provided"})
		return
	}
	err := Unsubscribe(address)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, nil)
}

func HandlerGetSubscribedAccounts(c *gin.Context) {
	c.IndentedJSON(http.StatusAccepted, gin.H{"accounts": GetSubscribedAccounts()})
}

func HandlerGetTransactions(c *gin.Context) {
	address := extractAddressFromPayload(c)
	if address == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "address not provided"})
		return
	}
	transactions, err := GetTransactions(address)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"transactions": transactions})
}

func HandlerGetAllTransactions(c *gin.Context) {
	transactions, err := GetAllTransactions()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusAccepted, gin.H{"transactions": transactions})
}

func extractAddressFromPayload(c *gin.Context) string {
	if address, ok := c.GetQuery("address"); ok {
		return address
	}
	return ""
}
