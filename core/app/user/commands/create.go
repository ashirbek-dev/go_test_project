package commands

import (
	"encoding/json"
	"fmt"
	"gateway/core/app_errors"
	"gateway/core/context"
	"gateway/core/data_types"
	"gateway/core/domain/entities"
	"gateway/core/domain/repositories"
	"gateway/core/domain/value_objects"
	"gateway/core/enums"
	"github.com/google/uuid"
	"strings"
	"time"
)

type CreateUserCommand struct {
	Pinfl          string    `json:"pinfl"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	PassportSerial string    `json:"passport_serial"`
	PassportNumber string    `json:"passport_number"`
	Phone          *string   `json:"phone"`
	WithSearch     bool      `json:"with_search"`
	BankId         uuid.UUID `json:"bank_id"`
}

type CreateUserResult struct {
	RefId string `json:"ref_id"`
}

type SearchCardsByPinflTask struct {
	UserId    uuid.UUID `json:"user_id"`
	Pinfl     string    `json:"pinfl"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

type SearchCardsByPhoneTask struct {
	UserId    uuid.UUID `json:"user_id"`
	Phone     string    `json:"phone"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

type CheckIntendUsersTask struct {
	UserRefId string `json:"user_ref_id"`
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
		handler.app.Logger.Error("user_create:from_json", err)
		return nil, err
	}
	return res, nil
}

func (handler _CreateUserHandler) Handle(command CreateUserCommand) (*CreateUserResult, error) {

	var err error
	var entity *entities.User
	var entityInfo *value_objects.UserInfo

	entity, err = handler.repo.GetUserByPinfl(command.Pinfl)

	if err != nil {
		handler.app.Logger.Error("user_create:get_by_pinfl", err)
		return nil, err
	}

	firstName := strings.ToUpper(command.FirstName)
	lastName := strings.ToUpper(command.LastName)
	passportSerial := strings.ToUpper(command.PassportSerial)
	passportNumber := strings.ToUpper(command.PassportNumber)

	if entity == nil {

		entity = &entities.User{
			Id:        uuid.New(),
			Pinfl:     command.Pinfl,
			CreatedAt: time.Now(),
		}

		err = handler.repo.CreateUser(*entity)

		if err != nil {
			handler.app.Logger.Error("user_create:create", err)
			return nil, err
		}

		b := uuid.New()
		refId := fmt.Sprintf("%X%X%X%X%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

		entityInfo = &value_objects.UserInfo{
			Id:             uuid.New(),
			RefId:          refId,
			UserId:         entity.Id,
			FirstName:      firstName,
			LastName:       lastName,
			PassportSerial: passportSerial,
			PassportNumber: passportNumber,
			CreatedAt:      time.Now(),
		}

		err = handler.repo.CreateUserInfo(*entityInfo)

		if err != nil {
			handler.app.Logger.Error("user_create:create_info", err)
			return nil, err
		}

	} else {

		entityInfo, err = handler.repo.GetUserInfo(entity.Id, passportSerial, passportNumber, firstName, lastName)

		if err != nil {
			handler.app.Logger.Error("user_create:get_user_info", err)
			return nil, err
		}

		if entityInfo == nil {

			b := uuid.New()
			refId := fmt.Sprintf("%X%X%X%X%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

			entityInfo = &value_objects.UserInfo{
				Id:             uuid.New(),
				RefId:          refId,
				UserId:         entity.Id,
				FirstName:      firstName,
				LastName:       lastName,
				PassportSerial: passportSerial,
				PassportNumber: passportNumber,
				CreatedAt:      time.Now(),
			}

			err = handler.repo.CreateUserInfo(*entityInfo)

			if err != nil {
				handler.app.Logger.Error("user_create:create_info", err)
				return nil, err
			}
		}
	}

	_ = handler.addPhone(entityInfo, command.Phone)

	/*if command.WithSearch {
		task := SearchCardsByPinflTask{
			UserId:    entityInfo.UserId,
			Pinfl:     entity.Pinfl,
			FirstName: entityInfo.FirstName,
			LastName:  entityInfo.LastName,
		}
		var taskData []byte
		taskData, err = json.Marshal(task)
		if err != nil {
			handler.app.Logger.Error("user_create:marshall_task", err)
		}
		pinflLast := entity.Pinfl[len(entity.Pinfl)-2:]
		qName := fmt.Sprintf("%s:%s", enums.QSearchCardsByPinflTask, pinflLast)
		handler.app.Kv.Push(qName, string(taskData))
	}*/

	intendTask := CheckIntendUsersTask{
		UserRefId: entityInfo.RefId,
	}
	var intendTaskData []byte
	intendTaskData, err = json.Marshal(intendTask)
	if err != nil {
		handler.app.Logger.Error("user_create:marshall_intend_task", err)
	}
	handler.app.Kv.Push(enums.QCheckKnownUsersTask, string(intendTaskData))

	return &CreateUserResult{
		RefId: entityInfo.RefId,
	}, nil
}

func (handler _CreateUserHandler) addPhone(userInfo *value_objects.UserInfo, phoneValue *string) error {
	var err error

	if userInfo == nil || phoneValue == nil || *phoneValue == "" {
		return nil
	}

	var phone *data_types.Phone

	phone, err = data_types.CreatePhone(*phoneValue)

	if err != nil {
		err = app_errors.NewAppErr(app_errors.INVALID_FORMAT, fmt.Sprintf("Phone: %s", phoneValue))
		handler.app.Logger.Error("user_create:parse_phone", err)
		return err
	}

	var userPhone *value_objects.UserPhone

	userPhone, err = handler.repo.GetUserPhone(userInfo.UserId, phone.PhoneNumber)

	if err != nil {
		handler.app.Logger.Error("user_create:get_by_ref_id", err)
		return err
	}

	if userPhone == nil {

		userPhone = &value_objects.UserPhone{
			Id:         uuid.New(),
			UserId:     userInfo.UserId,
			UserInfoId: userInfo.Id,
			Phone:      phone.PhoneNumber,
			CreatedAt:  time.Now(),
		}

		err = handler.repo.AddUserPhone(*userPhone)

		if err != nil {
			handler.app.Logger.Error("user_create:create", err)
			return err
		}

		task := SearchCardsByPhoneTask{
			UserId:    userInfo.UserId,
			Phone:     phone.PhoneNumber,
			FirstName: userInfo.FirstName,
			LastName:  userInfo.LastName,
		}
		var taskData []byte
		taskData, err = json.Marshal(task)
		if err != nil {
			handler.app.Logger.Error("user_create:marshall_task", err)
		}
		/*pinflLast := entity.Pinfl[len(entity.Pinfl)-2:]
		qName := fmt.Sprintf("%s:%s", enums.QSearchCardsByPinflTask, pinflLast)
		handler.app.Kv.Push(qName, string(taskData))*/
		handler.app.Kv.Push(enums.QSearchCardsByPhoneTask, string(taskData))
	}

	return nil
}
