package data

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/hossein1376/kala/internal/structure"
	"github.com/hossein1376/kala/internal/transfer"
)

type ProductsTable struct {
	DB *sqlx.DB
}

func (p *ProductsTable) Create(ctx context.Context, product *structure.Product) error {
	query := `
	INSERT INTO products (
	                      name,
	                      description,
	                      review,
	                      price,
	                      quantity
	                	)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id;`

	args, id := []any{product.Name, product.Description, product.Review, product.Price, product.Quantity}, 0
	err := p.DB.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		return err
	}

	product.ID = id
	return nil
}

func (p *ProductsTable) GetByID(ctx context.Context, id int) (*structure.Product, error) {
	query := `
	SELECT id, name, description, review, rating, rating_count, price, quantity
	FROM products
	WHERE id = $1 AND status = true;`

	var product structure.Product
	err := p.DB.QueryRowContext(ctx, query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Review,
		&product.Rating,
		&product.RatingCount,
		&product.Price,
		&product.Quantity,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, transfer.NotFoundError{}
		default:
			return nil, err
		}
	}

	return &product, nil
}

func (p *ProductsTable) GetAll(ctx context.Context) ([]structure.Product, error) {
	query := `
	SELECT id, name, description, review, rating, rating_count, price, quantity
	FROM products
	WHERE status = true
	ORDER BY id;`

	products := make([]structure.Product, 0)
	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product structure.Product
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Review,
			&product.Rating,
			&product.RatingCount,
			&product.Price,
			&product.Quantity,
		)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (p *ProductsTable) UpdateByID(ctx context.Context, id int, product *structure.Product) error {
	query := `
	UPDATE products
	Set
	    name = $2,
	    price = $3,
	    description = $4,
	    review = $5,
	    rating = $6,
	    rating_count = $7,
	    quantity = $8
	WHERE id = $1 AND status = true;`

	args := []any{
		id,
		product.Name,
		product.Price,
		product.Description,
		product.Review,
		product.Rating,
		product.RatingCount,
		product.Quantity,
	}
	result, err := p.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return transfer.NotFoundError{}
	}

	return nil
}

func (p *ProductsTable) DeleteByID(ctx context.Context, id int) error {
	query := `
	UPDATE products
	Set
	    status = false,
	    updated_at = now()
	WHERE id = $1 AND status = true;`

	result, err := p.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return transfer.NotFoundError{}
	}

	return nil
}
