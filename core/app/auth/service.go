package user

import (
	"gateway/core/app/user/commands"
	"gateway/core/app/user/queries"
	"gateway/core/context"
	"gateway/core/domain/repositories"
)

type Queries struct {
	GetUser queries.GetUserHandler
}

type Commands struct {
	CreateUser commands.CreateUserHandler
	AddPhone   commands.AddPhoneHandler
	PhoneFound commands.PhoneFoundHandler
}

type Service struct {
	Queries  Queries
	Commands Commands
}

func GetService(appCtx context.ApplicationContext, repo repositories.UserRepository) Service {
	return Service{
		Queries: Queries{
			GetUser: queries.NewGetUserHandler(appCtx, repo),
		},
		Commands: Commands{
			CreateUser: commands.GetCreateUserHandler(appCtx, repo),
			AddPhone:   commands.GetAddPhoneHandler(appCtx, repo),
			PhoneFound: commands.GetPhoneFoundHandler(appCtx, repo),
		},
	}
}
