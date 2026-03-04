package repositories

import (
	"gateway/core/domain/entities"
	"gateway/core/domain/value_objects"
	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserById(id uuid.UUID) (*entities.User, error)
	GetUserByPinfl(pinfl string) (*entities.User, error)
	CreateUser(user entities.User) error
	GetUserInfoByRefId(refId string) (*value_objects.UserInfo, error)
	CreateUserInfo(user value_objects.UserInfo) error
	GetUserInfo(userId uuid.UUID, passportSerial string, passportNumber string, firstName string, lastName string) (*value_objects.UserInfo, error)
	AddUserPhone(phone value_objects.UserPhone) error
	GetUserPhone(userId uuid.UUID, phone string) (*value_objects.UserPhone, error)
	GetUserInfosByUserId(userId uuid.UUID) ([]value_objects.UserInfo, error)
}
