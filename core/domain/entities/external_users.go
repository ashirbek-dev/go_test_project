package entities

import (
	"github.com/google/uuid"
	"time"
)

type ExternalUser struct {
	Id          uuid.UUID
	ServiceId   int64
	ExtUserId   int64
	Username    string
	Email       string
	Phone       string
	FirstName   string
	LastName    string
	IsActive    bool
	IsStaff     bool
	IsSuperUser bool
	ExtraData   map[string]interface{}
	LastLogin   time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
