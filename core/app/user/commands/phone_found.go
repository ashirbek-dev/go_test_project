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

type PhoneFoundCommand struct {
	UserId uuid.UUID `json:"user_id"`
	Phone  string    `json:"phone"`
}
type PhoneFoundHandler interface {
	Handle(command PhoneFoundCommand) error
	FromJson(jsonData []byte) (*PhoneFoundCommand, error)
}

type _PhoneFoundHandler struct {
	app  context.ApplicationContext
	repo repositories.UserRepository
}

func GetPhoneFoundHandler(app context.ApplicationContext, repo repositories.UserRepository) PhoneFoundHandler {
	return _PhoneFoundHandler{app: app, repo: repo}
}

func (handler _PhoneFoundHandler) FromJson(jsonData []byte) (*PhoneFoundCommand, error) {
	var res *PhoneFoundCommand
	err := json.Unmarshal(jsonData, &res)
	if err != nil {
		handler.app.Logger.Error("user_phone_found:from_json", err)
		return nil, err
	}
	return res, nil
}

func (handler _PhoneFoundHandler) Handle(command PhoneFoundCommand) error {
	var err error
	var userInfos []value_objects.UserInfo

	userInfos, err = handler.repo.GetUserInfosByUserId(command.UserId)

	if err != nil {
		handler.app.Logger.Error("user_phone_found:get_infos_by_user_id", err)
		return err
	}

	for _, userInfo := range userInfos {

		var phone *data_types.Phone

		phone, err = data_types.CreatePhone(command.Phone)

		if err != nil {
			err = app_errors.NewAppErr(app_errors.INVALID_FORMAT, fmt.Sprintf("Phone: %s", command.Phone))
			handler.app.Logger.Error("user_phone_found:parse_phone", err)
			continue
		}

		var userPhone *value_objects.UserPhone

		userPhone, err = handler.repo.GetUserPhone(userInfo.UserId, phone.PhoneNumber)

		if err != nil {
			handler.app.Logger.Error("user_phone_found:get_user_phone", err)
			continue
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
				handler.app.Logger.Error("user_phone_found:create_user_phone", err)
				continue
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
				handler.app.Logger.Error("user_phone_found:marshall_task", err)
			}
			handler.app.Kv.Push(enums.QSearchCardsByPhoneTask, string(taskData))
		}
	}

	return nil
}
