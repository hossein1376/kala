package data

import (
	"github.com/jmoiron/sqlx"

	"github.com/hossein1376/kala/internal/ent"
	"github.com/hossein1376/kala/internal/structure"
)

type Models struct {
	User    Users
	Product Products
}

func NewModels(client *sqlx.DB) *Models {
	return &Models{
		User:    &UsersTable{DB: client},
		Product: &ProductsTable{DB: client},
	}
}

type Users interface {
	Create(*structure.User) error
	GetByID(int) (*structure.User, error)
	GetByUsername(string) (*structure.User, error)
	GetAll(*structure.GetAllUsersRequest) ([]structure.User, int, error)
	UpdateByID(int, *structure.User) error
	DeleteByID(int) error
}

type Products interface {
	Create(product structure.Product) (*ent.Product, error)
	GetByID(id int) (*ent.Product, error)
	GetAll() ([]*ent.Product, error)
	UpdateByID(prod *ent.Product, id int) error
	DeleteByID(id int) error
}
