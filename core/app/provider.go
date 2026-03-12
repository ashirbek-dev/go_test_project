package app

import (
	"gateway/core/app/external_user"
	"gateway/core/app/token"
	"gateway/core/context"
	"gateway/core/domain/repositories"
)

type ApplicationService struct {
	Context context.ApplicationContext
}

func (app ApplicationService) GetTokenService(repo repositories.TokenRepository) token.Service {
	return token.GetService(app.Context, repo)
}

func (app ApplicationService) GetExtUserService(repo repositories.ExternalUserRepository) external_user.Service {
	return external_user.GetService(app.Context, repo)
}
