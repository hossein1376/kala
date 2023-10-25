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
	Create(structure.User) (*structure.User, error)
	GetByID(int) (*structure.User, error)
	GetByUsername(string) (*structure.User, error)
	GetAll(*structure.GetAllUsersRequest) ([]*structure.User, int, error)
	UpdateByID(int, *structure.User) error
	DeleteByID(int) error
}

type Product interface {
	Create(product structure.Product) (*ent.Product, error)
	GetByID(id int) (*ent.Product, error)
	GetAll() ([]*ent.Product, error)
	UpdateByID(prod *ent.Product, id int) error
	DeleteByID(id int) error
}
