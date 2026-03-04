package app

import (
	"gateway/core/app/user"
	"gateway/core/context"
	"gateway/core/domain/repositories"
)

type ApplicationService struct {
	Context context.ApplicationContext
}

func (app ApplicationService) GetUserService(repo repositories.UserRepository) user.Service {
	return user.GetService(app.Context, repo)
}
