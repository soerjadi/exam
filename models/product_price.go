package models

// ProductPrice model
type ProductPrice struct {
	ID        int64   `json:"id"`
	Amount    int64   `json:"amount"`
	Price     float64 `json:"price"`
	ProductID int64   `json:"product_id"`
}
