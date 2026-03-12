package queries

import (
	"gateway/core/context"
	"gateway/core/domain/repositories"
)

type ValidateTokenQuery struct {
	Token string
}

type ValidateTokenResponse struct {
	IsValid   bool
	ServiceId int64
}

type ValidateTokenHandler interface {
	Handle(query ValidateTokenQuery) (*ValidateTokenResponse, error)
}

type _ValidateTokenHandler struct {
	app  context.ApplicationContext
	repo repositories.TokenRepository
}

func NewValidateTokenHandler(app context.ApplicationContext, repo repositories.TokenRepository) ValidateTokenHandler {
	return _ValidateTokenHandler{app: app, repo: repo}
}

func (h _ValidateTokenHandler) Handle(query ValidateTokenQuery) (*ValidateTokenResponse, error) {
	isValid, serviceId, err := h.repo.IsValidToken(query.Token)

	if err != nil {
		return nil, err
	}

	return &ValidateTokenResponse{
		IsValid:   isValid,
		ServiceId: serviceId,
	}, nil
}
