package data

import (
	"github.com/hossein1376/kala/internal/ent"
	"github.com/hossein1376/kala/internal/structure"
)

type Models struct {
	User    User
	Product Product
	Image   Image
}

func NewModels(client *ent.Client) *Models {
	return &Models{
		User:    &UserModel{client: client},
		Product: &ProductModel{client: client},
		Image:   &ImageModel{client: client},
	}
}

type User interface {
	Create(user structure.User) (*ent.User, error)
	GetByID(id int) (*ent.User, error)
	GetByUsername(username string) (*ent.User, error)
	GetAll() ([]*ent.User, error)
	UpdateByID(id int, user *ent.User) error
	DeleteByID(id int) error
}

type Product interface {
	CreateNewProduct(product structure.Product) (*ent.Product, error)
	GetSingleProductByID(id int) (*ent.Product, error)
	GetAllProducts() ([]*ent.Product, error)
	UpdateProductByID(prod *ent.Product, id int) error
	DeleteProductByID(id int) error
}

type Image interface {
	Get(id ...int) ([]*ent.Image, error)
	Insert(image *structure.Image) error
}
