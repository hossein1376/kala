package structure

type Order struct {
	ID       int       `json:"id"`
	Products []Product `json:"products"`
	UserID   int       `json:"user_id"`
	Address  Address   `json:"address"`
}
