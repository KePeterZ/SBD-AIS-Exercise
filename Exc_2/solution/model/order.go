package model

type Order struct {
	DrinkID   uint64 `json:"drink_id"` // foreign key
	CreatedAt string `json:"created_at"`
	Amount    uint64 `json:"amount"`
}
