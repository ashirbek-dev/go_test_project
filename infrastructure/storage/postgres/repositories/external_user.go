package repositories

import (
	"database/sql"
	"errors"
	"gateway/core/domain/entities"
	"gateway/infrastructure/storage/postgres"
	"github.com/google/uuid"
)

type ExternalUser struct {
}

func (repo ExternalUser) GetUserById(id uuid.UUID) (*entities.ExternalUser, error) {
	var res entities.ExternalUser
	row := postgres.QueryRow(`
			SELECT
			    id, service_id, external_user_id,username,phone, email,first_name, last_name, is_active, is_staff,is_superuser,extra_data, last_login, created_at ,updated_at
			FROM external_users
			WHERE id = $1 `,
		id,
	)
	err := row.Scan(
		&res.Id,
		&res.ServiceId,
		&res.ExtUserId,
		&res.Username,
		&res.Phone,
		&res.Email,
		&res.FirstName,
		&res.LastName,
		&res.IsActive,
		&res.IsStaff,
		&res.IsSuperUser,
		&res.ExtraData,
		&res.LastLogin,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &res, nil
}

func (repo ExternalUser) GetUserByExtId(extId int64) (*entities.ExternalUser, error) {
	var res entities.ExternalUser
	row := postgres.QueryRow(`
			SELECT
			    id, service_id, external_user_id,username,phone, email,first_name, last_name, is_active, is_staff,is_superuser,extra_data, last_login, created_at ,updated_at
			FROM external_users
			WHERE external_user_id = $1`,
		extId,
	)
	err := row.Scan(
		&res.Id,
		&res.ServiceId,
		&res.ExtUserId,
		&res.Username,
		&res.Phone,
		&res.Email,
		&res.FirstName,
		&res.LastName,
		&res.IsActive,
		&res.IsStaff,
		&res.IsSuperUser,
		&res.ExtraData,
		&res.LastLogin,
		&res.CreatedAt,
		&res.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &res, nil
}

func (repo ExternalUser) CreateUser(m entities.ExternalUser) error {
	_, err := postgres.Exec(`
			INSERT INTO external_users (id, service_id, external_user_id,username,phone, email,first_name, last_name, is_active, is_staff,is_superuser,extra_data, last_login, created_at ,updated_at)
			VALUES ($1, $2, $3,$4,$5,$6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`,
		m.Id,
		m.ServiceId,
		m.ExtUserId,
		m.Username,
		m.Phone,
		m.Email,
		m.FirstName,
		m.LastName,
		m.IsActive,
		m.IsStaff,
		m.IsSuperUser,
		m.ExtraData,
		m.LastLogin,
		m.CreatedAt,
		m.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo ExternalUser) UpdateUser(m entities.ExternalUser) error {
	_, err := postgres.Exec(`
			UPDATE external_users
			SET username = $2,
			    email = $3,
			    phone = $4,
			    first_name = $5,
			    last_name = $6,
			    is_active = $7,
			    is_staff = $8,
			    is_superuser = $9,
			    extra_data = $10,
			    last_login = $11,
			    created_at  = $12,
				updated_at = $13
			WHERE id = $1 `,
		m.Username,
		m.Email,
		m.Phone,
		m.FirstName,
		m.LastName,
		m.IsActive,
		m.IsStaff,
		m.IsSuperUser,
		m.ExtraData,
		m.LastLogin,
		m.CreatedAt,
		m.UpdatedAt,
		m.Id)
	if err != nil {
		return err
	}
	return nil
}

func (repo ExternalUser) CheckUserByExternalId(externalId int64, serviceId int64) (bool, error) {
	var id string

	row := postgres.QueryRow(`
		SELECT id
		FROM external_users
		WHERE external_user_id = $1 AND service_id = $2
		LIMIT 1`,
		externalId, serviceId,
	)

	err := row.Scan(&id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
