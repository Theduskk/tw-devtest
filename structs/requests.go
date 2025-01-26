package structs

type EthGetLogsRequest struct {
	Id      int                `json:"id"`
	Jsonrpc string             `json:"jsonrpc"`
	Method  string             `json:"method"`
	Params  []EthGetLogsParams `json:"params"`
}

type EthGetLogsParams struct {
	FromBlock string   `json:"fromBlock"`
	ToBlock   string   `json:"toBlock"`
	Address   string   `json:"address"`
	Topics    []string `json:"topics"`
}
