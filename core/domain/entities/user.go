package entities

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID
	Pinfl     string
	CreatedAt time.Time
}
