package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"gateway/core/app_errors"
	"gateway/infrastructure/api/nats_connection"
	"os"
	"reflect"
	"time"
)

var Logger NatsLogger

type NatsLogger struct {
	natsService *nats_connection.Service
	subject     string
}

func InitLogger(subject string) {
	serviceUrl := os.Getenv("NATS_CLIENT_URL")
	streamName := os.Getenv("NATS_LOG_STREAM_NAME")

	natsService := nats_connection.Service{
		ServiceUrl: serviceUrl,
		StreamName: streamName,
	}

	Logger = NatsLogger{natsService: &natsService, subject: subject}
}

func (n NatsLogger) Log(action string, details ...any) {
	data := map[string]any{
		"action": action,
		"data":   details,
		"time":   time.Now().UnixMilli(),
	}
	payload, _ := json.Marshal(data)
	n.natsService.Send(n.subject, payload)
}

func (n NatsLogger) LogRequestResponse(action string, url string, request any, requestTime time.Time, responseCode int, response any) {
	now := time.Now()
	data := map[string]any{
		"action": action,
		"data": map[string]any{
			"url":                url,
			"request":            request,
			"request_time":       requestTime.Format("02.01.2006 15:04:05"),
			"request_timestamp":  requestTime.UnixMilli(),
			"response_code":      responseCode,
			"response":           response,
			"response_time":      now.Format("02.01.2006 15:04:05"),
			"response_timestamp": now.UnixMilli(),
		},
		"time": time.Now().UnixMilli(),
	}
	payload, _ := json.Marshal(data)
	n.natsService.Send(fmt.Sprintf(`%s_req_res`, n.subject), payload)
}

func (n NatsLogger) Error(action string, err error, details ...any) {
	var code int
	if errors.Is(err, &app_errors.ApplicationError{}) {
		if reflect.ValueOf(err).Kind() == reflect.Ptr {
			code = err.(*app_errors.ApplicationError).Code()
		} else {
			code = err.(app_errors.ApplicationError).Code()
		}
	}
	data := map[string]any{
		"action": action,
		"data": map[string]any{
			"error":   err.Error(),
			"code":    code,
			"details": details,
		},
		"time": time.Now().UnixMilli(),
	}
	payload, _ := json.Marshal(data)
	n.natsService.Send(fmt.Sprintf(`%s_error`, n.subject), payload)
}
