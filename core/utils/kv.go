package utils

import "time"

type KV interface {
	Push(key string, data any)
	PushTask(key string, data any) error
	Pop(key string) string
	SetEx(key string, value string, duration time.Duration)
	Get(key string) string
	LLen(key string) int64
}
