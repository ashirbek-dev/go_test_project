package commands

import (
	"encoding/json"
	"errors"
	"gateway/core/context"
	"gateway/core/domain/entities"
	"gateway/core/domain/repositories"
	"github.com/google/uuid"
)

type DeleteUserCommand struct {
	Id uuid.UUID `json:"id"`
}

type DeleteUserResult struct {
	Id string `json:"id"`
}

type DeleteUserHandler interface {
	Handle(command DeleteUserCommand) (*DeleteUserResult, error)
	FromJson(jsonData []byte) (*DeleteUserCommand, error)
}

type _DeleteUserHandler struct {
	app  context.ApplicationContext
	repo repositories.UserRepository
}

func GetDeleteUserHandler(app context.ApplicationContext, repo repositories.UserRepository) DeleteUserHandler {
	return _DeleteUserHandler{app: app, repo: repo}
}

func (handler _DeleteUserHandler) FromJson(jsonData []byte) (*DeleteUserCommand, error) {
	var res *DeleteUserCommand
	err := json.Unmarshal(jsonData, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (handler _DeleteUserHandler) Handle(command DeleteUserCommand) (*DeleteUserResult, error) {

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

	err = handler.repo.DeleteUser(command.Id)

	if err != nil {
		return nil, err
	}

	return &DeleteUserResult{
		Id: entity.Id.String(),
	}, nil
}
