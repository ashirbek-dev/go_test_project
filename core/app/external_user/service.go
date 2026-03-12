package external_user

import (
	"gateway/core/app/external_user/commands"
	"gateway/core/context"
	"gateway/core/domain/repositories"
)

type Commands struct {
	StoreExtUser commands.StoreUserHandler
}

type Service struct {
	Commands Commands
}

func GetService(appCtx context.ApplicationContext, repo repositories.ExternalUserRepository) Service {
	return Service{
		Commands: Commands{
			StoreExtUser: commands.GetStoreUserHandler(appCtx, repo),
		},
	}
}
