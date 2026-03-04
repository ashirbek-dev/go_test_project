package repositories

import (
	"gateway/core/domain/entities"
	"gateway/core/domain/value_objects"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserById(id uuid.UUID) (*entities.User, error)
	CreateUser(user entities.User) error
	UpdateUser(user entities.User) error
	DeleteUser(id uuid.UUID) error
	PaginateUsers(page int, pageSize int) ([]entities.User, error)
}
