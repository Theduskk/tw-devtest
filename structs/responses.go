package structs

// Structs for https://ethereum-rpc.publicnode.com/ api calls
type EthBlocknumberResponse struct {
	Id      int    `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  string `json:"result"`
}

type EthGetLogsResponse struct {
	Id      int           `json:"id"`
	Jsonrpc string        `json:"jsonrpc"`
	Result  []Transaction `json:"result"`
}
