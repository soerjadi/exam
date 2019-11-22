package models

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

// Product model
type Product struct {
	ID      int64     `json:"id"`
	Name    string    `json:"name"`
	SKU     string    `json:"SKU"`
	Created time.Time `json:"created"`
	Updated null.Time `json:"updated"`
}
