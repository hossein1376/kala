package structure

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Review      *string `json:"review"`
	Quantity    int32   `json:"quantity"`
	Price       int32   `json:"price"`
	Status      bool    `json:"status"`
	Available   bool    `json:"available"`
	Rating      float64 `json:"rating"`
	RatingCount int32   `json:"rating_count"`
}

type ProductUpdate struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Quantity    *int32   `json:"quantity"`
	Price       *int32   `json:"price"`
	Status      *bool    `json:"status"`
	Available   *bool    `json:"available"`
	Rating      *float64 `json:"rating"`
	RatingCount *int32   `json:"rating_count"`
}
