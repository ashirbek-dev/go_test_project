package repositories

import (
	"fmt"
	"gateway/infrastructure/storage/postgres"
)

type Token struct {
}

func (repo Token) IsValidToken(token string) (bool, int64, error) {
	var serviceId int64
	row := postgres.QueryRow("SELECT service_id FROM tokens WHERE token = $1 LIMIT 1", token)
	err := row.Scan(&serviceId)
	if err != nil {
		return false, 0, nil
	}
	fmt.Println("ServiceId:", serviceId)
	return true, serviceId, nil
}
