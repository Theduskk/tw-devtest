package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"tw-devtask/structs"

	"github.com/gin-gonic/gin"
)

type Parser interface {
	// add address to observer
	Subscribe(address string) error
	// remove address from observer
	Unsubscribe(address string) error
	// list of subscribed addresses
	GetSubscribedAccounts() []string
	// last parsed block
	GetCurrentBlock(address string) (string, error)
	// list of inbound or outbound transactions for an address
	GetTransactions(address string) ([]structs.Transaction, error)
	// list of inbound or outbound transactions for all subscribed addresses
	GetTransactionsForSubscribedAccounts() ([]structs.Transaction, error)
}

const endpoint string = "https://ethereum-rpc.publicnode.com/"
const jsonRpcVersion string = "2.0"

var subscribedAddresses []string

func Subscribe(address string) error {
	for _, subscribedAddress := range subscribedAddresses {
		if subscribedAddress == address {
			return gin.Error{Err: fmt.Errorf("address already subscribed"), Type: gin.ErrorTypePublic}
		}
	}
	subscribedAddresses = append(subscribedAddresses, address)
	fmt.Printf("Subscribed to address: %s", address)
	writePersistentData()
	return nil
}

func Unsubscribe(address string) error {
	for i, subscribedAddress := range subscribedAddresses {
		if subscribedAddress == address {
			subscribedAddresses = append(subscribedAddresses[:i], subscribedAddresses[i+1:]...)
			return nil
		}
	}
	writePersistentData()
	return gin.Error{Err: fmt.Errorf("address not found"), Type: gin.ErrorTypePublic}
}

func GetSubscribedAccounts() []string {
	readPersistentData()
	return subscribedAddresses
}

func GetCurrentBlock(address string) (string, error) {
	accountDetails, err := getDetailsForAccount(address)
	if err != nil {
		return "", err
	}
	if accountDetails.Result == nil || len(accountDetails.Result) == 0 {
		return "", fmt.Errorf("no transactions found for address %s", address)
	}
	return accountDetails.Result[0].BlockNumber, nil
}

func GetTransactions(address string) ([]structs.Transaction, error) {
	accountDetails, err := getDetailsForAccount(address)
	if err != nil {
		return nil, err
	}
	if accountDetails.Result == nil || len(accountDetails.Result) == 0 {
		return nil, fmt.Errorf("no transactions found for address %s", address)
	}
	return accountDetails.Result, nil
}

func GetAllTransactions() ([]structs.Transaction, error) {
	var transactions []structs.Transaction
	err := error(nil)
	readPersistentData()
	for _, address := range subscribedAddresses {
		accountDetails, err := getDetailsForAccount(address)
		if err != nil {
			continue
		}
		transactions = append(transactions, accountDetails.Result...)
	}
	return transactions, err
}

func getDetailsForAccount(account string) (*structs.EthGetLogsResponse, error) {
	jsonStr, err := json.Marshal(structs.EthGetLogsRequest{
		Id:      1,
		Jsonrpc: jsonRpcVersion,
		Method:  "eth_getLogs",
		Params: []structs.EthGetLogsParams{
			{
				FromBlock: "null",
				ToBlock:   "null",
				Address:   account,
				Topics:    []string{},
			},
		}})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}
	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, fmt.Errorf("failed to create POST request: %v", err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(request)
	fmt.Printf("Sent request to endpoint %s", endpoint)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var responseWrapper structs.EthGetLogsResponse
	if err := json.Unmarshal(body, &responseWrapper); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %v", err)
	}
	return &responseWrapper, nil
}

// Persists through restarts for go
// Does not persist through restarts for docker
// Also doesn't have proper error handling as this would most likely be removed for something like redis (For in memory)\
// or a database
func writePersistentData() {
	file, err := os.Create("subscribedAddresses")
	if err != nil {
		fmt.Printf("Failed to create file: %v", err)
		return
	}
	defer file.Close()
	for _, address := range subscribedAddresses {
		file.WriteString(address + "\n")
	}
	fmt.Printf("Wrote %v subscribed addresses to file", len(subscribedAddresses))
}

func readPersistentData() {
	file, err := os.Open("subscribedAddresses")
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
		return
	}
	defer file.Close()
	subscribedAddresses = nil
	for {
		var address string
		_, err := fmt.Fscanln(file, &address)
		if err != nil {
			break
		}
		subscribedAddresses = append(subscribedAddresses, address)
	}
	fmt.Printf("Read %v subscribed addresses from file", len(subscribedAddresses))
}
