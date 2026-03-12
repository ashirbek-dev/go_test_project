package repositories

import (
	"gateway/core/domain/entities"
	"github.com/google/uuid"
)

type ExternalUserRepository interface {
	GetUserById(id uuid.UUID) (*entities.ExternalUser, error)
	GetUserByExternalId(id int64, serviceId int64) (*entities.ExternalUser, error)
	CreateUser(user entities.ExternalUser) error
	UpdateUser(user entities.ExternalUser) error
	CheckUserByExternalId(externalId int64, serviceId int64) (bool, error)
}
