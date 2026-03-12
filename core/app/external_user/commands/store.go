package commands

import (
	"encoding/json"
	_ "errors"
	"gateway/core/context"
	"gateway/core/domain/entities"
	"gateway/core/domain/repositories"
	"github.com/google/uuid"
	"time"
)

type StoreUserCommand struct {
	ServiceId      int64                  `json:"service_id"`
	ExternalUserId int64                  `json:"external_user_id"`
	Username       string                 `json:"username"`
	Email          string                 `json:"email"`
	Phone          string                 `json:"phone"`
	FirstName      string                 `json:"first_name"`
	LastName       string                 `json:"last_name"`
	IsActive       bool                   `json:"is_active"`
	IsStaff        bool                   `json:"is_staff"`
	IsSuperUser    bool                   `json:"is_super_user"`
	ExtraData      map[string]interface{} `json:"extra_data"`
	LastLogin      time.Time              `json:"last_login"`
	CreatedAt      time.Time              `json:"created_at"`
	UpdatedAt      time.Time              `json:"updated_at"`
}

type StoreUserResult struct {
	Id string `json:"id"`
}

type StoreUserHandler interface {
	Handle(command StoreUserCommand) (*StoreUserResult, error)
	FromJson(jsonData []byte) (*StoreUserCommand, error)
}

type _StoreUserHandler struct {
	app  context.ApplicationContext
	repo repositories.ExternalUserRepository
}

func GetStoreUserHandler(app context.ApplicationContext, repo repositories.ExternalUserRepository) StoreUserHandler {
	return _StoreUserHandler{app: app, repo: repo}
}

func (handler _StoreUserHandler) FromJson(jsonData []byte) (*StoreUserCommand, error) {
	var res *StoreUserCommand
	err := json.Unmarshal(jsonData, &res)
	if err != nil {
		//handler.app.Logger.Error("user_create:from_json", err)
		return nil, err
	}
	return res, nil
}

func (handler _StoreUserHandler) Handle(command StoreUserCommand) (*StoreUserResult, error) {

	var err error
	var entity *entities.ExternalUser
	var exists bool

	exists, err = handler.repo.CheckUserByExternalId(command.ExternalUserId, command.ServiceId)

	if err != nil {
		//handler.app.Logger.Error("user_create:get_by_name", err)
		return nil, err
	}
	if exists {
		entity, err = handler.repo.GetUserByExternalId(command.ExternalUserId, command.ServiceId)
		if err != nil {
			return nil, err
		}
		entity.Username = command.Username
		entity.Email = command.Email
		entity.Phone = command.Phone
		entity.FirstName = command.FirstName
		entity.LastName = command.LastName
		entity.IsActive = command.IsActive
		entity.IsStaff = command.IsStaff
		entity.IsSuperUser = command.IsSuperUser
		entity.ExtraData = command.ExtraData
		entity.LastLogin = command.LastLogin
		entity.CreatedAt = command.CreatedAt
		entity.UpdatedAt = command.UpdatedAt

		err = handler.repo.UpdateUser(*entity)
		return nil, err
	}
	entity = &entities.ExternalUser{
		Id:          uuid.New(),
		ServiceId:   command.ServiceId,
		ExtUserId:   command.ExternalUserId,
		Username:    command.Username,
		Email:       command.Email,
		Phone:       command.Phone,
		FirstName:   command.FirstName,
		LastName:    command.LastName,
		IsActive:    command.IsActive,
		IsStaff:     command.IsStaff,
		IsSuperUser: command.IsSuperUser,
		ExtraData:   command.ExtraData,
		LastLogin:   command.LastLogin,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = handler.repo.CreateUser(*entity)

	if err != nil {
		//handler.app.Logger.Error("user_create:create", err)
		return nil, err
	}

	return &StoreUserResult{
		Id: entity.Id.String(),
	}, nil
}
