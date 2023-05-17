package data

import (
	"context"

	"kala/internal/ent"
	entUser "kala/internal/ent/user"
	"kala/internal/structure"
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
