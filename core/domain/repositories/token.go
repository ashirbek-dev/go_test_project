package repositories

type TokenRepository interface {
	IsValidToken(token string) (bool, int64, error)
}
