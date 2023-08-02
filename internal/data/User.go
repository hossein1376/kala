package data

import (
	"context"

	"github.com/hossein1376/kala/internal/ent"
	entUser "github.com/hossein1376/kala/internal/ent/user"
	"github.com/hossein1376/kala/internal/structure"
)

type UserModel struct {
	client *ent.Client
}

func (u *UserModel) CreateNewUser(user structure.User) error {
	_, err := u.client.User.Create().
		SetUsername(user.Username).
		SetPassword([]byte(user.Password.Hash)).
		SetEmail(user.Email).
		SetFirstName(user.FirstName).
		SetLastName(user.LastName).
		SetPhone(user.Phone).
		SetRole(entUser.RoleUser).
		Save(context.Background())

	return err
}

func (u *UserModel) GetSingleUserByID(id int) (*ent.User, error) {
	return u.client.User.Get(context.Background(), id)
}

func (u *UserModel) GetAllUsers() ([]*ent.User, error) {
	return u.client.User.Query().All(context.Background())
}

func (u *UserModel) UpdateUserByID(id int, user *ent.User) error {
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

func (u *UserModel) DeleteUserByID(id int) error {
	_, err := u.client.User.UpdateOneID(id).
		SetStatus(false).
		Save(context.Background())
	return err
}
