package context

import (
	"gateway/core/utils"
)

type ApplicationContext struct {
	Logger    utils.Logger
	Kv        utils.KV
	CryptoKey string
}
