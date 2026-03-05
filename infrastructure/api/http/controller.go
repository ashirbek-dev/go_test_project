package http

import (
	"encoding/json"
	"errors"
	"gateway/core/app"
	"gateway/core/app_errors"
	"gateway/infrastructure/api/http/controllers"
	"gateway/infrastructure/api/http/data_objects"
	"gateway/infrastructure/api/http/json_rpc_errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"reflect"
)

type Controller struct {
	appService *app.ApplicationService
}

func (n *Controller) PostHandler(c *gin.Context) {

	result, err := n.Handle(c)

	if err != nil {
		c.JSON(200, err)
		return
	}

	c.JSON(200, result)
}

func (n *Controller) GetHandler(c *gin.Context) {

	result, err := n.Handle(c)

	if err != nil {
		c.JSON(200, err)
		return
	}

	c.JSON(200, result)
}

type ControllerAction func(appSrv app.ApplicationService, payload []byte) (any, error)

var router map[string]ControllerAction

func getRouter() map[string]ControllerAction {

	if router == nil {
		userCtrl := controllers.UserController{}

		router = map[string]ControllerAction{
			"user.create": userCtrl.Create,
			"user.get":    userCtrl.Get,
		}
	}

	return router
}

func (n *Controller) Handle(c *gin.Context) (*data_objects.JsonRpcSuccessResponse, *data_objects.JsonRpcErrorResponse) {
	var jsonRpcRequest *data_objects.JsonRpcRequest

	// read request
	requestBytes, _readErr := ioutil.ReadAll(c.Request.Body)

	//ctx := context.RequestContext{}

	if _readErr != nil {
		//n.appService.Context.Logger.Error("handle_payload_data:reading", _readErr)
		return nil, data_objects.CreateErrorResponse(json_rpc_errors.CreateJsonRpcError(nil, json_rpc_errors.E_JSON_RPC_PARSER_ERROR, errors.New("invalid request")))
	}
	// validate json
	if !json.Valid(requestBytes) {
		//n.appService.Context.Logger.Error("handle_payload_data:validation", errors.New("invalid request"))
		return nil, data_objects.CreateErrorResponse(json_rpc_errors.CreateJsonRpcError(nil, json_rpc_errors.E_JSON_RPC_INVALID_REQUEST, errors.New("invalid request")))
	}

	parseErr := json.Unmarshal(requestBytes, &jsonRpcRequest)
	if parseErr != nil {
		//n.appService.Context.Logger.Error("handle_payload_data:unmarshal", parseErr)
		return nil, data_objects.CreateErrorResponse(json_rpc_errors.CreateJsonRpcError(nil, json_rpc_errors.E_JSON_RPC_INVALID_REQUEST, errors.New("invalid request")))
	}

	r := getRouter()
	handler, found := r[jsonRpcRequest.Method]

	if !found {
		//n.appService.Context.Logger.Error("handle_payload_data:method_not_found", errors.New("method not found"))
		return nil, data_objects.CreateErrorResponse(json_rpc_errors.CreateJsonRpcError(jsonRpcRequest.Id, json_rpc_errors.E_JSON_RPC_METHOD_NOT_FOUND, errors.New("method not found")))
	}

	var err error
	var marshal []byte
	marshal, err = json.Marshal(jsonRpcRequest.Params)
	if err != nil {
		//n.appService.Context.Logger.Error("handle_payload_data:marshal", err)
		return nil, data_objects.CreateErrorResponse(json_rpc_errors.CreateJsonRpcError(nil, json_rpc_errors.E_JSON_RPC_INVALID_REQUEST, errors.New("invalid request")))
	}

	var result any

	result, err = handler(*n.appService, marshal)

	if err != nil {
		var code int
		if errors.Is(err, &app_errors.ApplicationError{}) {
			if reflect.ValueOf(err).Kind() == reflect.Ptr {
				code = err.(*app_errors.ApplicationError).Code()
			} else {
				code = err.(app_errors.ApplicationError).Code()
			}
		}
		return nil, data_objects.CreateErrorResponse(json_rpc_errors.CreateJsonRpcError(jsonRpcRequest.Id, json_rpc_errors.E_JSON_RPC_OPERATION_ERROR+code, err))
	}

	return data_objects.CreateSuccessResponse(*jsonRpcRequest, result), nil
}
