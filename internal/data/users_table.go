package data

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/hossein1376/kala/internal/structure"
	"github.com/hossein1376/kala/internal/transfer"
)

var (
	duplicateUsernameError = `pq: duplicate key value violates unique constraint "users_username_key"`
)

type UsersTable struct {
	DB *sqlx.DB
}

func (u *UsersTable) Create(ctx context.Context, data *structure.User) error {
	query := `
	INSERT INTO users (
	                   username,
	                   password,
	                   first_name,
	                   last_name,
	                   email,
	                   phone
	                )
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id;`

	args, id := []any{data.Username, data.Password.Hash, data.FirstName, data.LastName, data.Email, data.Phone}, 0
	err := u.DB.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		switch {
		case err.Error() == duplicateUsernameError:
			return transfer.BadRequestError{}
		default:
			return err
		}
	}

	data.ID = id
	return nil
}

func (u *UsersTable) GetByID(ctx context.Context, id int) (*structure.User, error) {
	query := `
	SELECT id, username, password, first_name, last_name, email, phone, status
	FROM users
	WHERE id = $1;`

	var user structure.User
	err := u.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Password.Hash,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.Status,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, transfer.NotFoundError{}
		default:
			return nil, err
		}
	}

	if !user.Status {
		return nil, transfer.ForbiddenError{}
	}

	return &user, nil
}

func (u *UsersTable) GetByUsername(ctx context.Context, username string) (*structure.User, error) {
	query := `
	SELECT id, username, password, first_name, last_name, email, phone, status
	FROM users
	WHERE id = $1;`

	var user structure.User
	err := u.DB.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password.Hash,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Phone,
		&user.Status,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, transfer.NotFoundError{}
		default:
			return nil, err
		}
	}

	if !user.Status {
		return nil, transfer.ForbiddenError{}
	}

	return &user, nil
}

func (u *UsersTable) GetAll(ctx context.Context, pagination *structure.GetAllUsersRequest) ([]structure.User, int, error) {
	query := `
	SELECT id, username, password, first_name, last_name, email, phone
	FROM users
	WHERE status = true
	ORDER BY id
	LIMIT $1
	OFFSET $2;`

	users := make([]structure.User, 0, pagination.Count)
	rows, err := u.DB.QueryContext(ctx, query, pagination.Count, (pagination.Page-1)*pagination.Count)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var user structure.User

		err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password.Hash,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Phone,
		)
		if err != nil {
			return nil, 0, err
		}

		users = append(users, user)
	}

	query = `
	SELECT count(id)
	FROM users
	WHERE status = true;`

	var count int
	rows, err = u.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return nil, 0, err
		}
	}

	return users, count, nil
}

func (u *UsersTable) UpdateByID(ctx context.Context, id int, user *structure.User) error {
	query := `
	UPDATE users
	SET
	    username = $2,
	    password = $3,
	    first_name = $4,
	    last_name = $5,
	    email = $6,
	    phone = $7,
	    updated_at = now()
	WHERE id = $1 AND status = true AND role != 'admin';`

	args := []any{id, user.Username, user.Password.Hash, user.FirstName, user.LastName, user.Email, user.Phone}
	result, err := u.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return transfer.NotFoundError{}
	}

	return nil
}

func (u *UsersTable) DeleteByID(ctx context.Context, id int) error {
	query := `
	UPDATE users
	Set
	    status = false,
	    updated_at = now()
	WHERE id = $1 AND status = true AND role != 'admin';`

	result, err := u.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return transfer.NotFoundError{}
	}

	return nil
}
