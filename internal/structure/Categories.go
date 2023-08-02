package structure

type Category struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Image       Image         `json:"image"`
	SubCategory []SubCategory `json:"subcategory"`
	Product     []Product     `json:"product"`
	Brand       []Brand       `json:"brand"`
}

type SubCategory struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Image       Image      `json:"image"`
	Category    []Category `json:"category"`
	Product     []Product  `json:"product"`
}
