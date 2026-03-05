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

type CreateUserCommand struct {
	Name string `json:"name"`
}

type CreateUserResult struct {
	Id string `json:"id"`
}

type CreateUserHandler interface {
	Handle(command CreateUserCommand) (*CreateUserResult, error)
	FromJson(jsonData []byte) (*CreateUserCommand, error)
}

type _CreateUserHandler struct {
	app  context.ApplicationContext
	repo repositories.UserRepository
}

func GetCreateUserHandler(app context.ApplicationContext, repo repositories.UserRepository) CreateUserHandler {
	return _CreateUserHandler{app: app, repo: repo}
}

func (handler _CreateUserHandler) FromJson(jsonData []byte) (*CreateUserCommand, error) {
	var res *CreateUserCommand
	err := json.Unmarshal(jsonData, &res)
	if err != nil {
		//handler.app.Logger.Error("user_create:from_json", err)
		return nil, err
	}
	return res, nil
}

func (handler _CreateUserHandler) Handle(command CreateUserCommand) (*CreateUserResult, error) {

	var err error
	var entity *entities.User

	entity, err = handler.repo.GetUserByName(command.Name)

	if err != nil {
		//handler.app.Logger.Error("user_create:get_by_name", err)
		return nil, err
	}
	if entity != nil {
		return nil, errors.New("user already exists this name")
	}
	entity = &entities.User{
		Id:        uuid.New(),
		Name:      command.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = handler.repo.CreateUser(*entity)

	if err != nil {
		//handler.app.Logger.Error("user_create:create", err)
		return nil, err
	}

	return &CreateUserResult{
		Id: entity.Id.String(),
	}, nil
}
