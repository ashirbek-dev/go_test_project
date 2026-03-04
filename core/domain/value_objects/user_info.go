package value_objects

import (
	"github.com/google/uuid"
	"time"
)

type UserInfo struct {
	Id             uuid.UUID
	RefId          string
	UserId         uuid.UUID
	FirstName      string
	LastName       string
	PassportSerial string
	PassportNumber string
	CreatedAt      time.Time
}
