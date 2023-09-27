package data

import (
	"context"
	"time"

	"github.com/hossein1376/kala/internal/ent"
	entProduct "github.com/hossein1376/kala/internal/ent/product"
	"github.com/hossein1376/kala/internal/structure"
)

type ProductModel struct {
	client *ent.Client
}

func (p *ProductModel) CreateNewProduct(product structure.Product) (*ent.Product, error) {
	return p.client.Product.Create().
		SetName(product.Name).
		SetPrice(product.Price).
		SetStatus(product.Status).
		SetAvailable(product.Available).
		SetDescription(product.Description).
		SetRating(product.Rating).
		SetRatingCount(product.RatingCount).
		SetQuantity(product.Quantity).
		SetCreateTime(time.Now()).
		SetUpdateTime(time.Now()).
		Save(context.Background())
}

func (p *ProductModel) GetSingleProductByID(id int) (*ent.Product, error) {
	return p.client.Product.Query().
		Where(entProduct.ID(id)).
		Only(context.Background())
}

func (p *ProductModel) GetAllProducts() ([]*ent.Product, error) {
	return p.client.Product.Query().
		All(context.Background())
}

func (p *ProductModel) UpdateProductByID(prod *ent.Product, id int) error {
	_, err := p.client.Product.UpdateOneID(id).
		Where(entProduct.Available(true), entProduct.Status(true)).
		SetName(prod.Name).
		SetPrice(prod.Price).
		SetStatus(prod.Status).
		SetAvailable(prod.Available).
		SetDescription(prod.Description).
		SetRating(prod.Rating).
		SetRatingCount(prod.RatingCount).
		SetQuantity(prod.Quantity).
		SetUpdateTime(time.Now()).
		Save(context.Background())

	return err
}

func (p *ProductModel) DeleteProductByID(id int) error {
	_, err := p.client.Product.UpdateOneID(id).
		Where(entProduct.Available(true), entProduct.Status(true)).
		SetStatus(false).
		SetUpdateTime(time.Now()).
		Save(context.Background())
	return err
}
