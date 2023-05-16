package data

import (
	"context"

	"kala/internal/ent"
)

type UserModel struct {
	client *ent.Client
}

func (u *UserModel) GetSingleUserByID(id int) (*ent.User, error) {
	return u.client.User.Get(context.Background(), id)
}

func (u *UserModel) GetAllUsers() ([]*ent.User, error) {
	return u.client.User.Query().All(context.Background())
}
