package commands

import (
	"encoding/json"
	"errors"
	"gateway/core/context"
	"gateway/core/domain/entities"
	"gateway/core/domain/repositories"
	"github.com/google/uuid"
	"time"
)

type UpdateUserCommand struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type UpdateUserResult struct {
	Id string `json:"id"`
}

type UpdateUserHandler interface {
	Handle(command UpdateUserCommand) (*UpdateUserResult, error)
	FromJson(jsonData []byte) (*UpdateUserCommand, error)
}

type _UpdateUserHandler struct {
	app  context.ApplicationContext
	repo repositories.UserRepository
}

func GetUpdateUserHandler(app context.ApplicationContext, repo repositories.UserRepository) UpdateUserHandler {
	return _UpdateUserHandler{app: app, repo: repo}
}

func (handler _UpdateUserHandler) FromJson(jsonData []byte) (*UpdateUserCommand, error) {
	var res *UpdateUserCommand
	err := json.Unmarshal(jsonData, &res)
	if err != nil {
		//handler.app.Logger.Error("user_update:from_json", err)
		return nil, err
	}
	return res, nil
}

func (handler _UpdateUserHandler) Handle(command UpdateUserCommand) (*UpdateUserResult, error) {
	var err error
	var entity *entities.User

	entity, err = handler.repo.GetUserById(command.Id)

	if err != nil {
		//handler.app.Logger.Error("user_create:get_by_name", err)
		return nil, err
	}
	if entity == nil {
		return nil, errors.New("entity not found")
	}

	entity = &entities.User{
		Name:      command.Name,
		UpdatedAt: time.Now(),
	}

	err = handler.repo.UpdateUser(*entity)

	if err != nil {
		//handler.app.Logger.Error("user_create:create", err)
		return nil, err
	}

	return &UpdateUserResult{
		Id: entity.Id.String(),
	}, nil
}
