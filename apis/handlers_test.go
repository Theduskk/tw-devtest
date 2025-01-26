package apis

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"tw-devtask/structs"

	gin "github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

var mocks mock.Mock

func Setup() *gin.Engine {
	router := gin.Default()
	mocks = mock.Mock{}

	return router
}

var getLogsResponse = structs.EthGetLogsResponse{
	Id:      1,
	Jsonrpc: "2.0",
	Result: []structs.Transaction{
		{
			BlockHash:   "0x",
			BlockNumber: "0x",
		},
	},
}

func TestHandlerSubscribe(t *testing.T) {

	router := Setup()
	router.POST("/subscribe", HandlerSubscribe)
	request, _ := http.NewRequest("POST", "/subscribe?address=bonk", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, response.Code)
	}
}
func TestHandlerUnsubscribe(t *testing.T) {

	router := Setup()
	router.POST("/unsubscribe", HandlerUnsubscribe)
	request, _ := http.NewRequest("POST", "/unsubscribe?address=bonk", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, response.Code)
	}
}
func TestHandlerGetSubscribedAccounts(t *testing.T) {
	router := Setup()
	router.GET("/getSubscribedAccounts", HandlerGetSubscribedAccounts)
	request, _ := http.NewRequest("GET", "/getSubscribedAccounts", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, response.Code)
	}
}

func TestHandlerGetCurrentBlock(t *testing.T) {
	router := Setup()
	mocks.On("getDetailsForAccount", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2").Return(getLogsResponse, nil)

	router.GET("/getCurrentBlock", HandlerGetCurrentBlock)
	request, _ := http.NewRequest("GET", "/getCurrentBlock?address=0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, response.Code)
	}

}

// Write a test for HandlerGetTransactions
func TestHandlerGetTransactions(t *testing.T) {
	router := Setup()
	mocks.On("getDetailsForAccount", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2").Return(getLogsResponse, nil)

	router.GET("/getTransactions", HandlerGetTransactions)
	request, _ := http.NewRequest("GET", "/getTransactions?address=0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, response.Code)
	}
}

// Write a test for HandlerGetAllTransactions
func TestHandlerGetAllTransactions(t *testing.T) {
	router := Setup()
	mocks.On("getDetailsForAccount", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2").Return(getLogsResponse, nil)

	router.GET("/getAllTransactions", HandlerGetAllTransactions)
	request, _ := http.NewRequest("GET", "/getAllTransactions", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, response.Code)
	}
}

// Write a test for HandlerGetCurrentBlock that returns an error
func TestHandlerGetCurrentBlockError(t *testing.T) {
	router := Setup()
	mocks.On("getDetailsForAccount", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2").Return(nil, fmt.Errorf("error"))

	router.GET("/getCurrentBlock", HandlerGetCurrentBlock)
	request, _ := http.NewRequest("GET", "/getCurrentBlock?address=0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, response.Code)
	}

}

// Write a test for HandlerGetTransactions that returns an error
func TestHandlerGetTransactionsError(t *testing.T) {
	router := Setup()
	mocks.On("getDetailsForAccount", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2").Return(nil, fmt.Errorf("error"))

	router.GET("/getTransactions", HandlerGetTransactions)
	request, _ := http.NewRequest("GET", "/getTransactions?address=0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, response.Code)
	}
}

// Write a test for HandlerGetAllTransactions that returns an error
func TestHandlerGetAllTransactionsError(t *testing.T) {
	router := Setup()
	mocks.On("getDetailsForAccount", "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2").Return(nil, fmt.Errorf("error"))

	router.GET("/getAllTransactions", HandlerGetAllTransactions)
	request, _ := http.NewRequest("GET", "/getAllTransactions", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, response.Code)
	}
}

// Write a test for HandlerSubscribe that returns an error
func TestHandlerSubscribeError(t *testing.T) {
	router := Setup()
	router.POST("/subscribe", HandlerSubscribe)
	request, _ := http.NewRequest("POST", "/subscribe", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, response.Code)
	}
}

// Write a test for HandlerUnsubscribe that returns an error
func TestHandlerUnsubscribeError(t *testing.T) {
	router := Setup()
	router.POST("/unsubscribe", HandlerUnsubscribe)
	request, _ := http.NewRequest("POST", "/unsubscribe", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, response.Code)
	}
}

// Write a test for HandlerGetSubscribedAccounts that returns an error
func TestHandlerGetSubscribedAccountsError(t *testing.T) {
	router := Setup()
	router.GET("/getSubscribedAccounts", HandlerGetSubscribedAccounts)
	request, _ := http.NewRequest("GET", "/getSubscribedAccounts", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, response.Code)
	}
}
