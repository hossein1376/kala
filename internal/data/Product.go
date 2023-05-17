package data

import (
	"context"

	"kala/internal/ent"
	"kala/internal/ent/comment"
	"kala/internal/ent/product"
	"kala/internal/structure"
)

type ProductModel struct {
	client *ent.Client
}

func (p *ProductModel) CreateNewProduct(product structure.Product) (*ent.Product, error) {
	prod, err := p.client.Product.Create().
		SetName(product.Name).
		SetPrice(product.Price).
		Save(context.Background())
	return prod, err
}

func (p *ProductModel) GetSingleProductByID(id int) (*ent.Product, error) {
	return p.client.Product.Query().
		Where(product.ID(id)).
		WithBrand().
		WithComment(func(query *ent.CommentQuery) {
			query.
				Where(comment.StatusEQ(comment.StatusPublished)).
				Order(ent.Asc(comment.FieldID))
		}).
		WithValues(func(query *ent.AttributeValueQuery) {
			query.
				WithAttributes()
		}).
		WithImage().
		Only(context.Background())
}

func (p *ProductModel) GetAllProducts() ([]*ent.Product, error) {
	return p.client.Product.Query().
		WithBrand().
		WithComment(func(query *ent.CommentQuery) {
			query.
				Where(comment.StatusEQ(comment.StatusPublished)).
				Order(ent.Asc(comment.FieldID))
		}).
		WithValues(func(query *ent.AttributeValueQuery) {
			query.
				WithAttributes()
		}).
		WithImage().
		All(context.Background())
}

func (p *ProductModel) DeleteProductByID(id int) error {
	_, err := p.client.Product.UpdateOneID(id).
		SetStatus(false).
		Save(context.Background())
	return err
}
