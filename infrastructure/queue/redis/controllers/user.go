package controllers

import (
	"incard.uz/humo/core/app"
	"incard.uz/humo/infrastructure/storage/postgres/repositories"
)

type UserController struct {
}

func (c UserController) PhoneFound(appSrv app.ApplicationService, queueName string, payload []byte, eventAt int64) {
	srv := appSrv.GetUserService(repositories.User{})
	cmd, err := srv.Commands.PhoneFound.FromJson(payload)
	if err != nil {
		return
	}
	handleErr := srv.Commands.PhoneFound.Handle(*cmd)
	if handleErr != nil {
		return
	}
}

func (c UserController) AddPhone(appSrv app.ApplicationService, queueName string, payload []byte, eventAt int64) {
	srv := appSrv.GetUserService(repositories.User{})
	cmd, err := srv.Commands.AddPhone.FromJson(payload)
	if err != nil {
		return
	}
	_, handleErr := srv.Commands.AddPhone.Handle(*cmd)
	if handleErr != nil {
		return
	}
}
