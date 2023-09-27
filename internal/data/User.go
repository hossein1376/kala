package data

import (
	"context"
	"fmt"

	"github.com/hossein1376/kala/internal/ent"
	entUser "github.com/hossein1376/kala/internal/ent/user"
	"github.com/hossein1376/kala/internal/response"
	"github.com/hossein1376/kala/internal/structure"
	"github.com/hossein1376/kala/pkg/Password"
)

type UserModel struct {
	client *ent.Client
}

func (u *UserModel) Create(data structure.User) (*structure.User, error) {
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
			return nil, response.UserCreationError{Msg: err.Error()}
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

func (u *UserModel) GetByID(id int) (*structure.User, error) {
	user, err := u.client.User.Get(context.Background(), id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return nil, response.UserRetrievalError{Msg: fmt.Sprintf("user with id '%v' was not found", id)}
		default:
			return nil, err
		}
	}

	return &structure.User{
		ID:        user.ID,
		Username:  user.Username,
		Password:  Password.Password{Hash: user.Password},
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}, nil
}

func (u *UserModel) GetByUsername(username string) (*structure.User, error) {
	user, err := u.client.User.
		Query().
		Where(entUser.Username(username)).
		First(context.Background())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return nil, response.UserNotFound{Username: &username}
		default:
			return nil, err
		}
	}

	if !user.Status {
		return nil, response.UserDisabled{Username: &username}
	}

	return &structure.User{
		ID:        user.ID,
		Username:  user.Username,
		Password:  Password.Password{Hash: user.Password},
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}, err
}

func (u *UserModel) GetAll() ([]*structure.User, error) {
	users, err := u.client.User.Query().All(context.Background())
	if err != nil {
		return nil, err
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

	return resp, err
}

func (u *UserModel) UpdateByID(id int, user *structure.User) error {
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
			return response.UserUpdateError{Msg: fmt.Sprintf("failed to update user with id '%v'", id)}
		default:
			return err
		}
	}

	return nil
}

func (u *UserModel) DeleteByID(id int) error {
	_, err := u.client.User.UpdateOneID(id).
		SetStatus(false).
		Save(context.Background())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return response.UserDeleteError{Msg: fmt.Sprintf("failed to delete user with id '%v'", id)}
		default:
			return err
		}
	}

	return nil
}
