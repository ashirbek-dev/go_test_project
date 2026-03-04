package repositories

import (
	"database/sql"
	"errors"
	"gateway/core/domain/entities"
	"gateway/core/domain/value_objects"
	"gateway/infrastructure/storage/postgres"
	"github.com/google/uuid"
)

type User struct {
}

func (repo User) GetUserById(id uuid.UUID) (*entities.User, error) {
	var res entities.User
	row := postgres.QueryRow(`
			SELECT 
			    id, pinfl, created_at 
			FROM users 
			WHERE id = $1`,
		id,
	)
	err := row.Scan(
		&res.Id,
		&res.Pinfl,
		&res.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &res, nil
}
func (repo User) GetUserByPinfl(pinfl string) (*entities.User, error) {
	var res entities.User
	row := postgres.QueryRow(`
			SELECT 
			    id, pinfl, created_at 
			FROM users 
			WHERE pinfl = $1`,
		pinfl,
	)
	err := row.Scan(
		&res.Id,
		&res.Pinfl,
		&res.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &res, nil
}
func (repo User) CreateUser(m entities.User) error {
	_, err := postgres.Exec(`
			INSERT INTO users (id, pinfl, created_at) 
			VALUES ($1, $2, $3)`,
		m.Id,
		m.Pinfl,
		m.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
func (repo User) GetUserInfoByRefId(refId string) (*value_objects.UserInfo, error) {
	var res value_objects.UserInfo
	row := postgres.QueryRow(`
			SELECT 
			    id, ref_id, user_id, f_name, l_name, passport_serial, passport_number, created_at
			FROM user_infos 
			WHERE ref_id = $1`,
		refId,
	)
	err := row.Scan(
		&res.Id,
		&res.RefId,
		&res.UserId,
		&res.FirstName,
		&res.LastName,
		&res.PassportSerial,
		&res.PassportNumber,
		&res.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &res, nil
}
func (repo User) CreateUserInfo(m value_objects.UserInfo) error {
	_, err := postgres.Exec(`
			INSERT INTO user_infos (id, ref_id, user_id, f_name, l_name, passport_serial, passport_number, created_at) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		m.Id,
		m.RefId,
		m.UserId,
		m.FirstName,
		m.LastName,
		m.PassportSerial,
		m.PassportNumber,
		m.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
func (repo User) GetUserInfo(userId uuid.UUID, passportSerial string, passportNumber string, firstName string, lastName string) (*value_objects.UserInfo, error) {
	var res value_objects.UserInfo
	row := postgres.QueryRow(`
			SELECT 
			   id, ref_id, user_id, f_name, l_name, passport_serial, passport_number, created_at
			FROM user_infos 
			WHERE user_id = $1 AND passport_serial = $2 AND passport_number = $3 AND f_name = $4 AND l_name = $5`,
		userId,
		passportSerial,
		passportNumber,
		firstName,
		lastName,
	)
	err := row.Scan(
		&res.Id,
		&res.RefId,
		&res.UserId,
		&res.FirstName,
		&res.LastName,
		&res.PassportSerial,
		&res.PassportNumber,
		&res.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &res, nil
}
func (repo User) AddUserPhone(m value_objects.UserPhone) error {
	_, err := postgres.Exec(`
			INSERT INTO user_phones (id, user_id, user_info_id, phone, created_at) 
			VALUES ($1, $2, $3, $4, $5)`,
		m.Id,
		m.UserId,
		m.UserInfoId,
		m.Phone,
		m.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
func (repo User) GetUserPhone(userId uuid.UUID, phone string) (*value_objects.UserPhone, error) {
	var res value_objects.UserPhone
	row := postgres.QueryRow(`
			SELECT 
			    id, user_id, user_info_id, phone, created_at
			FROM user_phones 
			WHERE user_id = $1 AND phone = $2`,
		userId,
		phone,
	)
	err := row.Scan(
		&res.Id,
		&res.UserId,
		&res.UserInfoId,
		&res.Phone,
		&res.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &res, nil
}
func (repo User) GetUserInfosByUserId(userId uuid.UUID) ([]value_objects.UserInfo, error) {
	var res []value_objects.UserInfo
	rows, err := postgres.Query(`
			SELECT 
			    id, ref_id, user_id, f_name, l_name, passport_serial, passport_number, created_at
			FROM user_infos 
			WHERE user_id = $1`,
		userId,
	)
	if err != nil {
		return res, err
	}
	defer rows.Close()
	for rows.Next() {
		var item value_objects.UserInfo
		if err = rows.Scan(
			&item.Id,
			&item.RefId,
			&item.UserId,
			&item.FirstName,
			&item.LastName,
			&item.PassportSerial,
			&item.PassportNumber,
			&item.CreatedAt,
		); err != nil {
			return res, err
		}

		res = append(res, item)
	}
	if err = rows.Err(); err != nil {
		return res, err
	}
	return res, nil
}
