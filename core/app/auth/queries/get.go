package queries

import (
	"encoding/json"
	"fmt"
	"gateway/core/app_errors"
	"gateway/core/context"
	"gateway/core/domain/entities"
	"gateway/core/domain/repositories"
	"gateway/core/domain/value_objects"
	"strings"
)

type GetUserQuery struct {
	RefId string `json:"ref_id"`
}

type GetUserQueryResult struct {
	Pinfl          string `json:"pinfl"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	PassportSerial string `json:"passport_serial"`
	PassportNumber string `json:"passport_number"`
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
		handler.app.Logger.Error("user_get:from_json", err)
		return nil, err
	}
	return res, nil
}

func (handler _GetUserHandler) Handle(query GetUserQuery) (*GetUserQueryResult, error) {

	var err error
	var entity *entities.User
	var entityInfo *value_objects.UserInfo

	entityInfo, err = handler.repo.GetUserInfoByRefId(strings.ToUpper(query.RefId))

	if err != nil {
		handler.app.Logger.Error("user_get:get_by_ref_id", err)
		return nil, err
	}

	if entityInfo == nil {
		err = app_errors.NewAppErr(app_errors.CLIENT_NOT_FOUND, fmt.Sprintf("ID: %s", query.RefId))
		handler.app.Logger.Error("user_get:get_by_ref_id", err)
		return nil, err
	}

	entity, err = handler.repo.GetUserById(entityInfo.UserId)

	if err != nil {
		handler.app.Logger.Error("user_get:get_by_ref_id", err)
		return nil, err
	}

	return &GetUserQueryResult{
		Pinfl:          entity.Pinfl,
		FirstName:      entityInfo.FirstName,
		LastName:       entityInfo.LastName,
		PassportSerial: entityInfo.PassportSerial,
		PassportNumber: entityInfo.PassportNumber,
	}, nil

}
