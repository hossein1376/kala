package data

import (
	"github.com/hossein1376/kala/internal/ent"
	"github.com/hossein1376/kala/internal/structure"
)

type Models struct {
	User    User
	Product Product
}

func NewModels(client *ent.Client) *Models {
	return &Models{
		User:    &UserModel{client: client},
		Product: &ProductModel{client: client},
	}
}

type User interface {
	Create(user structure.User) (*structure.User, error)
	GetByID(id int) (*structure.User, error)
	GetByUsername(username string) (*structure.User, error)
	GetAll() ([]*structure.User, error)
	UpdateByID(id int, user *structure.User) error
	DeleteByID(id int) error
}

type Product interface {
	Create(product structure.Product) (*ent.Product, error)
	GetByID(id int) (*ent.Product, error)
	GetAll() ([]*ent.Product, error)
	UpdateByID(prod *ent.Product, id int) error
	DeleteByID(id int) error
}
