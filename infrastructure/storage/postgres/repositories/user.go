package repositories

import (
	"database/sql"
	"errors"
	"gateway/core/domain/entities"
	"gateway/infrastructure/storage/postgres"
	"github.com/google/uuid"
)

type User struct {
}

func (repo User) DeleteUser(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (repo User) GetUserById(id uuid.UUID) (*entities.User, error) {
	var res entities.User
	row := postgres.QueryRow(`
			SELECT 
			    id, name, created_at ,updated_at
			FROM users 
			WHERE id = $1`,
		id,
	)
	err := row.Scan(
		&res.Id,
		&res.Name,
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

func (repo User) GetUserByName(name string) (*entities.User, error) {
	var res entities.User
	row := postgres.QueryRow(`
			SELECT 
			    id, name, created_at,updated_at
			FROM users 
			WHERE name = $1`,
		name,
	)
	err := row.Scan(
		&res.Id,
		&res.Name,
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

func (repo User) CreateUser(m entities.User) error {
	_, err := postgres.Exec(`
			INSERT INTO users (id, name, created_at, updated_at) 
			VALUES ($1, $2, $3,$4)`,
		m.Id,
		m.Name,
		m.CreatedAt,
		m.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (repo User) UpdateUser(m entities.User) error {
	_, err := postgres.Exec(`
			UPDATE users
			SET name = $2, updated_at = $3
			WHERE id = $1 `,
		m.Name,
		m.UpdatedAt,
		m.Id)
	if err != nil {
		return err
	}
	return nil
}

func (repo User) DeleteUserById(id uuid.UUID) error {
	_, err := postgres.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo User) PaginateUsers(page int, limit int) ([]entities.User, int, error) {
	offset := (page - 1) * limit
	query, err := postgres.Query(`
			SELECT id, name, created_at,updated_at FROM users LIMIT $1 OFFSET $2
			`, offset, limit)

	if err != nil {
		return nil, 0, err
	}
	defer func(query *sql.Rows) {
		err := query.Close()
		if err != nil {
			return
		}
	}(query)

	var res []entities.User
	for query.Next() {
		var item entities.User
		err = query.Scan(&item.Id, &item.Name, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, 0, err
		}
		res = append(res, item)
	}

	return res, limit, nil
}
