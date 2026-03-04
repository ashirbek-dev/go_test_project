package data_objects

type JsonRpcRequest struct {
	Version *string `json:"jsonrpc"`
	Id      *string `json:"id"`
	Method  string  `json:"method"`
	Params  any     `json:"params"`
}
