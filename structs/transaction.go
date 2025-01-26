package structs

// Transaction struct
// Source https://ethereum.github.io/execution-apis/api-documentation
type Transaction struct {
	LogIndex         string   `json:"logIndex"`
	Removed          bool     `json:"removed"`
	BlockNumber      string   `json:"blockNumber"`
	BlockHash        string   `json:"blockHash"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
	Address          string   `json:"address"`
	Data             string   `json:"data"`
	Topics           []string `json:"topics"`
}
