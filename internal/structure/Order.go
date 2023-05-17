package structure

type Order struct {
	Products []Product `json:"products"`
	UserID   int       `json:"user_id"`
	Address  Address   `json:"address"`
}
