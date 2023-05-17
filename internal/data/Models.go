package data

import (
	"kala/internal/ent"
)

type Models struct {
	User    UserModel
	Product ProductModel
}

func NewModels(client *ent.Client) Models {
	return Models{
		User:    UserModel{client: client},
		Product: ProductModel{client: client},
	}
}
