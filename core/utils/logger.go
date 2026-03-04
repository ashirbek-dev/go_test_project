package utils

import "time"

type Logger interface {
	Log(action string, details ...any)
	Error(action string, err error, details ...any)
	LogRequestResponse(action string, url string, request any, requestTime time.Time, responseCode int, response any)
}
