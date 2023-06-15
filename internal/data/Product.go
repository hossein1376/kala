package data

import (
	"context"
	"time"

	"kala/internal/ent"
	"kala/internal/ent/comment"
	entProduct "kala/internal/ent/product"
	"kala/internal/structure"
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
		SetBrand(product.Brand).
		SetRating(product.Rating).
		SetRatingCount(product.RatingCount).
		SetQuantity(product.Quantity).
		AddImage(product.Image...).
		AddCategory(product.Category...).
		AddSubCategory(product.SubCategory...).
		SetCreateTime(time.Now()).
		SetUpdateTime(time.Now()).
		Save(context.Background())
}

func (p *ProductModel) GetSingleProductByID(id int) (*ent.Product, error) {
	return p.client.Product.Query().
		Where(entProduct.ID(id)).
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

func (p *ProductModel) UpdateProductByID(prod *ent.Product, id int) error {
	_, err := p.client.Product.UpdateOneID(id).
		Where(entProduct.Available(true), entProduct.Status(true)).
		SetName(prod.Name).
		SetPrice(prod.Price).
		SetStatus(prod.Status).
		SetAvailable(prod.Available).
		SetDescription(prod.Description).
		SetBrand(prod.Edges.Brand).
		SetRating(prod.Rating).
		SetRatingCount(prod.RatingCount).
		SetQuantity(prod.Quantity).
		AddImage(prod.Edges.Image...).
		AddCategory(prod.Edges.Category...).
		AddSubCategory(prod.Edges.SubCategory...).
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
