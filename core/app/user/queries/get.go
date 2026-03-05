package queries

import (
	"encoding/json"
	"errors"
	"gateway/core/context"
	"gateway/core/domain/entities"
	"gateway/core/domain/repositories"
	"github.com/google/uuid"
)

type GetUserQuery struct {
	Id uuid.UUID `json:"id"`
}

type GetUserQueryResult struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetUserHandler interface {
	Handle(command GetUserQuery) (*GetUserQueryResult, error)
	FromJson(jsonData []byte) (*GetUserQuery, error)
}

type _GetUserHandler struct {
	app  context.ApplicationContext
	repo repositories.UserRepository
}

func NewGetUserHandler(app context.ApplicationContext, repo repositories.UserRepository) GetUserHandler {
	return _GetUserHandler{app: app, repo: repo}
}

func (handler _GetUserHandler) FromJson(jsonData []byte) (*GetUserQuery, error) {
	var res *GetUserQuery
	err := json.Unmarshal(jsonData, &res)
	if err != nil {
		//handler.app.Logger.Error("user_get:from_json", err)
		return nil, err
	}
	return res, nil
}

func (handler _GetUserHandler) Handle(query GetUserQuery) (*GetUserQueryResult, error) {

	var err error
	var entity *entities.User

	entity, err = handler.repo.GetUserById(query.Id)

	if err != nil {
		return nil, err
	}

	if entity == nil {
		return nil, errors.New("entity not found")
	}

	return &GetUserQueryResult{
		Id:        entity.Id.String(),
		Name:      entity.Name,
		CreatedAt: entity.CreatedAt.String(),
		UpdatedAt: entity.UpdatedAt.String(),
	}, nil

}
