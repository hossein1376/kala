package data

import (
	"context"

	"github.com/hossein1376/kala/internal/ent"
	entUser "github.com/hossein1376/kala/internal/ent/user"
	"github.com/hossein1376/kala/internal/response"
	"github.com/hossein1376/kala/internal/structure"
)

type UserModel struct {
	client *ent.Client
}

func (u *UserModel) Create(user structure.User) (*ent.User, error) {
	return u.client.User.Create().
		SetUsername(user.Username).
		SetPassword([]byte(user.Password.Hash)).
		SetEmail(user.Email).
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetPhone(user.Phone).
		SetRole(entUser.RoleUser).
		Save(context.Background())
}

func (u *UserModel) GetByID(id int) (*ent.User, error) {
	return u.client.User.Get(context.Background(), id)
}

func (u *UserModel) GetByUsername(username string) (*ent.User, error) {
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

	return user, err
}

func (u *UserModel) GetAll() ([]*ent.User, error) {
	return u.client.User.Query().All(context.Background())
}

func (u *UserModel) UpdateByID(id int, user *ent.User) error {
	_, err := u.client.User.UpdateOneID(id).
		SetUsername(user.Username).
		SetPassword(user.Password).
		SetEmail(user.Email).
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetPhone(user.Phone).
		Save(context.Background())
	return err
}

func (u *UserModel) DeleteByID(id int) error {
	_, err := u.client.User.UpdateOneID(id).
		SetStatus(false).
		Save(context.Background())
	return err
}
