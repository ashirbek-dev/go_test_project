package repositories

import (
	"gateway/core/domain/entities"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserById(id uuid.UUID) (*entities.User, error)
	GetUserByName(name string) (*entities.User, error)
	CreateUser(user entities.User) error
	UpdateUser(user entities.User) error
	DeleteUser(id uuid.UUID) error
	PaginateUsers(page int, limit int) ([]entities.User, int, error)
}
