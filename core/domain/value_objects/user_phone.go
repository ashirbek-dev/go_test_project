package value_objects

import (
	"github.com/google/uuid"
	"time"
)

type UserPhone struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	UserInfoId uuid.UUID
	Phone      string
	CreatedAt  time.Time
}
