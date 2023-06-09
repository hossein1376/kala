package structure

import (
	"kala/internal/ent"
)

type Product struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Quantity    int32              `json:"quantity"`
	Price       int32              `json:"price"`
	Status      bool               `json:"status"`
	Available   bool               `json:"available"`
	Rating      float64            `json:"rating"`
	RatingCount int32              `json:"rating_count"`
	Brand       *ent.Brand         `json:"brand"`
	Category    []*ent.Category    `json:"category"`
	SubCategory []*ent.SubCategory `json:"sub_category"`
	Image       []*ent.Image       `json:"image"`
	Comment     []*ent.Comment     `json:"comment"`
}

type ProductUpdate struct {
	Name        *string            `json:"name"`
	Description *string            `json:"description"`
	Quantity    *int32             `json:"quantity"`
	Price       *int32             `json:"price"`
	Status      *bool              `json:"status"`
	Available   *bool              `json:"available"`
	Rating      *float64           `json:"rating"`
	RatingCount *int32             `json:"rating_count"`
	Brand       *ent.Brand         `json:"brand"`
	Category    []*ent.Category    `json:"category"`
	SubCategory []*ent.SubCategory `json:"sub_category"`
	Images      []*ent.Image       `json:"images"`
	Comments    []*ent.Comment     `json:"comments"`
}
