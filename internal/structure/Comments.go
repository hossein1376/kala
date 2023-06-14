package structure

type Comment struct {
	ID          int      `json:"id"`
	User        int      `json:"user_id"`
	Product     int      `json:"product_id"`
	Content     string   `json:"content"`
	Status      string   `json:"status"`
	Likes       int32    `json:"likes"`
	Dislikes    int32    `json:"dislikes"`
	Rating      float64  `json:"rating"`
	RotingCount int      `json:"roting_count"`
	Verified    bool     `json:"verified"`
	Pros        []string `json:"pros"`
	Cons        []string `json:"cons"`
	Image       []Image  `json:"image"`
}
