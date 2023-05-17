package structure

type Product struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Price       int32       `json:"price"`
	Rating      float64     `json:"rating"`
	RatingCount int32       `json:"rating_count"`
	Category    Category    `json:"category"`
	SubCategory SubCategory `json:"sub_category"`
	Comments    []Comment   `json:"comments"`
}
