package data

import (
	"kala/internal/ent"
	"kala/internal/structure"
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
	CreateNewUser(user structure.User) error
	GetSingleUserByID(id int) (*ent.User, error)
	GetAllUsers() ([]*ent.User, error)
	UpdateUserByID(id int, user *ent.User) error
	DeleteUserByID(id int) error
}

type Product interface {
	CreateNewProduct(product structure.Product) (*ent.Product, error)
	GetSingleProductByID(id int) (*ent.Product, error)
	GetAllProducts() ([]*ent.Product, error)
	UpdateProductByID(prod *ent.Product, id int) error
	DeleteProductByID(id int) error
}
