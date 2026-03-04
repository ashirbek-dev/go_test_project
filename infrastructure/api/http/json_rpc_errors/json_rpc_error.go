package json_rpc_errors

type JsonRpcError struct {
	Id   *string
	Code int
	Err  error
}

func (ins *JsonRpcError) Error() (*string, string, int) {
	return ins.Id, ins.Err.Error(), ins.Code
}

func CreateJsonRpcError(id *string, code int, err error) *JsonRpcError {
	return &JsonRpcError{Id: id, Code: code, Err: err}
}

const E_JSON_RPC_PARSER_ERROR = -32700
const E_JSON_RPC_INVALID_REQUEST = -32600
const E_JSON_RPC_METHOD_NOT_FOUND = -32601
const E_JSON_RPC_INVALID_PARAMS = -32602
const E_JSON_RPC_INTERNAL_ERROR = -32603
const E_JSON_RPC_OPERATION_ERROR = -33000
