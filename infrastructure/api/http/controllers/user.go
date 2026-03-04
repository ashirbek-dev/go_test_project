package controllers

import (
	"gateway/core/app"
	"gateway/infrastructure/storage/postgres/repositories"
)

type UserController struct {
}

func (c UserController) Create(appSrv app.ApplicationService, payload []byte) (any, error) {
	srv := appSrv.GetUserService(repositories.User{})
	cmd, err := srv.Commands.CreateUser.FromJson(payload)
	if err != nil {
		return nil, err
	}
	res, handleErr := srv.Commands.CreateUser.Handle(*cmd)
	if handleErr != nil {
		return nil, handleErr
	}
	return res, nil
}
func (c UserController) AddPhone(appSrv app.ApplicationService, payload []byte) (any, error) {
	srv := appSrv.GetUserService(repositories.User{})
	cmd, err := srv.Commands.AddPhone.FromJson(payload)
	if err != nil {
		return nil, err
	}
	res, handleErr := srv.Commands.AddPhone.Handle(*cmd)
	if handleErr != nil {
		return nil, handleErr
	}
	return res, nil
}

func (c UserController) Get(appSrv app.ApplicationService, payload []byte) (any, error) {
	srv := appSrv.GetUserService(repositories.User{})
	cmd, err := srv.Queries.GetUser.FromJson(payload)
	if err != nil {
		return nil, err
	}
	res, handleErr := srv.Queries.GetUser.Handle(*cmd)
	if handleErr != nil {
		return nil, handleErr
	}
	return res, nil
}
