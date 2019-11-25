package types

import "github.com/soerjadi/exam/models"

// Product represent product model with product category
type Product struct {
	ID       int64              `json:"id"`
	Name     string             `json:"name"`
	SKU      string             `json:"sku"`
	Category []*models.Category `json:"category"`
}
