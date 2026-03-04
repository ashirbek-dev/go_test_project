package value_objects

import (
	"github.com/google/uuid"
	"time"
)

type UserInfo struct {
	Id             uuid.UUID
	Name      string
	CreatedAt      time.Time
}
