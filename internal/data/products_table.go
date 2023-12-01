package data

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/hossein1376/kala/internal/ent"
	entProduct "github.com/hossein1376/kala/internal/ent/product"
	"github.com/hossein1376/kala/internal/structure"
)

type ProductsTable struct {
	client *ent.Client
	*sqlx.DB
}

func (p *ProductsTable) Create(product structure.Product) (*ent.Product, error) {
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

func (p *ProductsTable) GetByID(id int) (*ent.Product, error) {
	return p.client.Product.Query().
		Where(entProduct.ID(id)).
		Only(context.Background())
}

func (p *ProductsTable) GetAll() ([]*ent.Product, error) {
	return p.client.Product.Query().
		All(context.Background())
}

func (p *ProductsTable) UpdateByID(prod *ent.Product, id int) error {
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

func (p *ProductsTable) DeleteByID(id int) error {
	_, err := p.client.Product.UpdateOneID(id).
		Where(entProduct.Available(true), entProduct.Status(true)).
		SetStatus(false).
		SetUpdateTime(time.Now()).
		Save(context.Background())
	return err
}
