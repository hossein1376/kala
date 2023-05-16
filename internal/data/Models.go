package data

import (
	"kala/internal/ent"
)

type Models struct {
	User UserModel
}

func NewModels(client *ent.Client) Models {
	return Models{
		User: UserModel{client: client},
	}
}
