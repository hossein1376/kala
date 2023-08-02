package structure

import (
	"github.com/hossein1376/kala/internal/ent"
)

type Product struct {
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Quantity    int32           `json:"quantity"`
	Price       int32           `json:"price"`
	Status      bool            `json:"status"`
	Available   bool            `json:"available"`
	Rating      float64         `json:"rating"`
	RatingCount int32           `json:"rating_count"`
	Brand       *ent.Brand      `json:"brand"`
	Category    []*ent.Category `json:"category"`
	Image       []*ent.Image    `json:"image"`
}

type ProductUpdate struct {
	Name        *string         `json:"name"`
	Description *string         `json:"description"`
	Quantity    *int32          `json:"quantity"`
	Price       *int32          `json:"price"`
	Status      *bool           `json:"status"`
	Available   *bool           `json:"available"`
	Rating      *float64        `json:"rating"`
	RatingCount *int32          `json:"rating_count"`
	Brand       *ent.Brand      `json:"brand"`
	Category    []*ent.Category `json:"category"`
	Images      []*ent.Image    `json:"images"`
}
