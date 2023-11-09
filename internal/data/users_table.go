package data

import (
	"context"

	"github.com/hossein1376/kala/internal/ent"
	entUser "github.com/hossein1376/kala/internal/ent/user"
	"github.com/hossein1376/kala/internal/structure"
	"github.com/hossein1376/kala/internal/transfer"
	"github.com/hossein1376/kala/pkg/password"
)

type UsersTable struct {
	client *ent.Client
}

func (u *UsersTable) Create(data structure.User) (*structure.User, error) {
	user, err := u.client.User.Create().
		SetUsername(data.Username).
		SetPassword([]byte(data.Password.Hash)).
		SetEmail(data.Email).
		SetFirstName(data.FirstName).
		SetLastName(data.LastName).
		SetPhone(data.Phone).
		SetRole(entUser.RoleUser).
		Save(context.Background())
	if err != nil {
		switch {
		case ent.IsConstraintError(err) || ent.IsValidationError(err):
			return nil, transfer.BadRequestError{Err: err}
		default:
			return nil, err
		}
	}

	return &structure.User{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}, nil
}

func (u *UsersTable) GetByID(id int) (*structure.User, error) {
	user, err := u.client.User.
		Query().
		Where(
			entUser.ID(id),
			entUser.Status(true),
		).
		First(context.Background())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return nil, transfer.NotFoundError{}
		default:
			return nil, err
		}
	}

	return &structure.User{
		ID:        user.ID,
		Username:  user.Username,
		Password:  password.Password{Hash: user.Password},
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}, nil
}

func (u *UsersTable) GetByUsername(username string) (*structure.User, error) {
	user, err := u.client.User.
		Query().
		Where(entUser.Username(username)).
		First(context.Background())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return nil, transfer.NotFoundError{Err: err}
		default:
			return nil, err
		}
	}

	if !user.Status {
		return nil, transfer.ForbiddenError{}
	}

	return &structure.User{
		ID:        user.ID,
		Username:  user.Username,
		Password:  password.Password{Hash: user.Password},
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}, err
}

func (u *UsersTable) GetAll(query *structure.GetAllUsersRequest) ([]*structure.User, int, error) {
	users, err := u.client.User.Query().
		Where(
			entUser.Status(true),
		).
		Limit(query.Count).
		Offset((query.Page - 1) * query.Count).
		All(context.Background())
	if err != nil {
		return nil, 0, err
	}

	resp := make([]*structure.User, 0, len(users))
	for _, user := range users {
		r := &structure.User{
			ID:        user.ID,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
		}
		resp = append(resp, r)
	}

	count, err := u.client.User.Query().
		Where(
			entUser.Status(true),
		).
		Count(context.Background())

	return resp, count, err
}

func (u *UsersTable) UpdateByID(id int, user *structure.User) error {
	_, err := u.client.User.UpdateOneID(id).
		SetUsername(user.Username).
		SetPassword(user.Password.Hash).
		SetEmail(user.Email).
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetPhone(user.Phone).
		Save(context.Background())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return transfer.NotFoundError{Err: err}
		default:
			return err
		}
	}

	return nil
}

func (u *UsersTable) DeleteByID(id int) error {
	_, err := u.client.User.UpdateOneID(id).
		SetStatus(false).
		Save(context.Background())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return transfer.NotFoundError{Err: err}
		default:
			return err
		}
	}

	return nil
}
