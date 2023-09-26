package data

import (
	"context"
	"fmt"

	"github.com/hossein1376/kala/internal/ent"
	entUser "github.com/hossein1376/kala/internal/ent/user"
	"github.com/hossein1376/kala/internal/structure"
	est "github.com/hossein1376/kala/pkg/EST"
)

type UserModel struct {
	client *ent.Client
}

func (u *UserModel) Create(data structure.User) (*structure.UserResponse, error) {
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
			return nil, est.BadRequestError(err)
		default:
			return nil, est.InternalError(err)
		}
	}

	return &structure.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}, nil
}

func (u *UserModel) GetByID(id int) (*structure.UserResponse, error) {
	user, err := u.client.User.Get(context.Background(), id)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return nil, est.NotFoundError(fmt.Errorf("user with id '%d' was not found", id))
		default:
			return nil, est.InternalError(err)
		}
	}

	return &structure.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}, nil
}

func (u *UserModel) GetByUsername(username string) (*structure.UserResponse, error) {
	user, err := u.client.User.
		Query().
		Where(entUser.Username(username)).
		First(context.Background())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return nil, est.NotFoundError(fmt.Errorf("user with username '%s' was not found", username))
		default:
			return nil, est.InternalError(err)
		}
	}

	if !user.Status {
		return nil, est.ForbiddenError(nil)
	}

	return &structure.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
	}, nil
}

func (u *UserModel) GetAll() ([]*structure.UserResponse, error) {
	users, err := u.client.User.Query().All(context.Background())
	if err != nil {
		return nil, est.InternalError(err)
	}

	resp := make([]*structure.UserResponse, 0, len(users))
	for _, user := range users {
		r := &structure.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
		}
		resp = append(resp, r)
	}

	return resp, nil
}

func (u *UserModel) UpdateByID(id int, user *structure.UserResponse) error {
	_, err := u.client.User.UpdateOneID(id).
		SetUsername(user.Username).
		SetPassword(user.Password).
		SetEmail(user.Email).
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetPhone(user.Phone).
		Save(context.Background())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return est.NotFoundError(fmt.Errorf("failed to update user with id '%d', %d", id, err))
		case ent.IsConstraintError(err):
			return est.BadRequestError(fmt.Errorf("failed to update user with id '%d', %d", id, err))
		default:
			return est.InternalError(err)
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
			return est.NotFoundError(fmt.Errorf("failed to delete user with id '%d'", id))
		default:
			return est.InternalError(err)
		}
	}

	return nil
}
