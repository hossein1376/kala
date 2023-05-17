package structure

type Seller struct {
	Name         string     `json:"name"`
	Category     []Category `json:"category"`
	Product      []Product  `json:"product"`
	Rating       float64    `json:"rating"`
	RatingNumber int32      `json:"rating_number"`
}
