package commands

import (
	"encoding/json"
	"fmt"
	"gateway/core/app_errors"
	"gateway/core/context"
	"gateway/core/data_types"
	"gateway/core/domain/repositories"
	"gateway/core/domain/value_objects"
	"gateway/core/enums"
	"github.com/google/uuid"
	"time"
)

type AddPhoneCommand struct {
	RefId string `json:"ref_id"`
	Phone string `json:"phone"`
}

type AddPhoneResult struct {
	RefId string `json:"ref_id"`
	Phone string `json:"phone"`
}

type AddPhoneHandler interface {
	Handle(command AddPhoneCommand) (*AddPhoneResult, error)
	FromJson(jsonData []byte) (*AddPhoneCommand, error)
}

type _AddPhoneHandler struct {
	app  context.ApplicationContext
	repo repositories.UserRepository
}

func GetAddPhoneHandler(app context.ApplicationContext, repo repositories.UserRepository) AddPhoneHandler {
	return _AddPhoneHandler{app: app, repo: repo}
}

func (handler _AddPhoneHandler) FromJson(jsonData []byte) (*AddPhoneCommand, error) {
	var res *AddPhoneCommand
	err := json.Unmarshal(jsonData, &res)
	if err != nil {
		handler.app.Logger.Error("user_add_phone:from_json", err)
		return nil, err
	}
	return res, nil
}

func (handler _AddPhoneHandler) Handle(command AddPhoneCommand) (*AddPhoneResult, error) {
	var err error
	var entityInfo *value_objects.UserInfo

	entityInfo, err = handler.repo.GetUserInfoByRefId(command.RefId)

	if err != nil {
		handler.app.Logger.Error("user_add_phone:get_by_ref_id", err)
		return nil, err
	}

	if entityInfo == nil {
		err = app_errors.NewAppErr(app_errors.CLIENT_NOT_FOUND, fmt.Sprintf("ID: %s", command.RefId))
		handler.app.Logger.Error("user_add_phone:get_by_ref_id", err)
		return nil, err
	}

	var phone *data_types.Phone

	phone, err = data_types.CreatePhone(command.Phone)
	if err != nil {
		err = app_errors.NewAppErr(app_errors.INVALID_FORMAT, fmt.Sprintf("Phone: %s", command.Phone))
		handler.app.Logger.Error("user_add_phone:parse_phone", err)
		return nil, err
	}

	var userPhone *value_objects.UserPhone

	userPhone, err = handler.repo.GetUserPhone(entityInfo.UserId, phone.PhoneNumber)

	if err != nil {
		handler.app.Logger.Error("user_add_phone:get_by_ref_id", err)
		return nil, err
	}

	if userPhone == nil {

		userPhone = &value_objects.UserPhone{
			Id:         uuid.New(),
			UserId:     entityInfo.UserId,
			UserInfoId: entityInfo.Id,
			Phone:      phone.PhoneNumber,
			CreatedAt:  time.Now(),
		}

		err = handler.repo.AddUserPhone(*userPhone)

		if err != nil {
			handler.app.Logger.Error("user_add_phone:create", err)
			return nil, err
		}

		task := SearchCardsByPhoneTask{
			UserId:    entityInfo.UserId,
			Phone:     phone.PhoneNumber,
			FirstName: entityInfo.FirstName,
			LastName:  entityInfo.LastName,
		}
		var taskData []byte
		taskData, err = json.Marshal(task)
		if err != nil {
			handler.app.Logger.Error("user_create:marshall_task", err)
		}
		handler.app.Kv.Push(enums.QSearchCardsByPhoneTask, string(taskData))
	}

	return &AddPhoneResult{
		RefId: entityInfo.RefId,
		Phone: userPhone.Phone,
	}, nil
}
