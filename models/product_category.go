package models

type ProductCategory struct {
	ID         int64 `json:"id"`
	ProductID  int64 `json:"product_id"`
	CategoryID int64 `json:"category_id"`
}
