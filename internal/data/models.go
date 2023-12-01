package data

import (
	"github.com/jmoiron/sqlx"

	"github.com/hossein1376/kala/internal/structure"
)

type Models struct {
	User    Users
	Product Products
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{
		User:    &UsersTable{DB: db},
		Product: &ProductsTable{DB: db},
	}
}

type Users interface {
	Create(user *structure.User) error
	GetByID(id int) (*structure.User, error)
	GetByUsername(username string) (*structure.User, error)
	GetAll(pagination *structure.GetAllUsersRequest) ([]structure.User, int, error)
	UpdateByID(id int, user *structure.User) error
	DeleteByID(id int) error
}

type Products interface {
	Create(product *structure.Product) error
	GetByID(id int) (*structure.Product, error)
	GetAll() ([]structure.Product, error)
	UpdateByID(id int, product *structure.Product) error
	DeleteByID(id int) error
}
