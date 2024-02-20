package data

import (
	"context"

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
	Create(ctx context.Context, user *structure.User) error
	GetByID(ctx context.Context, id int) (*structure.User, error)
	GetByUsername(ctx context.Context, username string) (*structure.User, error)
	GetAll(ctx context.Context, pagination *structure.GetAllUsersRequest) ([]structure.User, int, error)
	UpdateByID(ctx context.Context, id int, user *structure.User) error
	DeleteByID(ctx context.Context, id int) error
}

type Products interface {
	Create(ctx context.Context, product *structure.Product) error
	GetByID(ctx context.Context, id int) (*structure.Product, error)
	GetAll(ctx context.Context) ([]structure.Product, error)
	UpdateByID(ctx context.Context, id int, product *structure.Product) error
	DeleteByID(ctx context.Context, id int) error
}
