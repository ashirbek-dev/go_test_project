package data_objects

import (
	"fmt"
	"gateway/infrastructure/api/http/json_rpc_errors"
	"github.com/google/uuid"
)

type JsonRpcResponse interface {
}

type JsonRpcErrorResponse struct {
	Version string `json:"jsonrpc"`
	Id      string `json:"id"`
	Error   any    `json:"error"`
}

type JsonRpcSuccessResponse struct {
	Version string `json:"jsonrpc"`
	Id      string `json:"id"`
	Result  any    `json:"result"`
}

func CreateErrorResponse(err *json_rpc_errors.JsonRpcError) *JsonRpcErrorResponse {
	instance := JsonRpcErrorResponse{}
	instance.Version = "2.0"
	if err.Id != nil {
		instance.Id = *err.Id
		if instance.Id == "" {
			instance.Id = createRequestId()
		}
	} else {
		instance.Id = createRequestId()
	}
	instance.Error = map[string]interface{}{
		"code":    err.Code,
		"message": err.Err.Error(),
	}
	return &instance
}

func CreateSuccessResponse(request JsonRpcRequest, result any) *JsonRpcSuccessResponse {
	instance := JsonRpcSuccessResponse{}
	instance.Version = "2.0"
	if request.Id != nil {
		instance.Id = *request.Id
		if instance.Id == "" {
			instance.Id = createRequestId()
		}
	} else {
		instance.Id = createRequestId()
	}
	instance.Result = result
	return &instance
}
func createRequestId() string {
	b := uuid.New()
	return fmt.Sprintf("%X%X%X%X%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
