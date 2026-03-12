package token

import (
	"gateway/core/app/token/queries"
	"gateway/core/context"
	"gateway/core/domain/repositories"
)

type Queries struct {
	ValidateToken queries.ValidateTokenHandler
}

type Service struct {
	Queries Queries
}

func GetService(appCtx context.ApplicationContext, repo repositories.TokenRepository) Service {
	return Service{
		Queries: Queries{
			ValidateToken: queries.NewValidateTokenHandler(appCtx, repo),
		},
	}
}
