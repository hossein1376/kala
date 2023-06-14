package structure

type Brand struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      bool       `json:"status"`
	Rating      float64    `json:"rating"`
	RotingCount int        `json:"roting_count"`
	Image       Image      `json:"image"`
	Category    []Category `json:"category"`
	Product     []Product  `json:"product"`
}
