package models

import "time"

// Order model
type Order struct {
	ID        int64     `json:"id"`
	ProductID int64     `json:"product_id"`
	Amount    int64     `json:"amount"`
	Price     float64   `json:"price"`
	Status    int       `json:"status"`
	Created   time.Time `json:"created"`
}

// OrderPending for initialize pending payment order
var OrderPending = 0

// OrderProccessed order that payment is verified and waiting to be shipped
var OrderProccessed = 1

// OrderShipped order proccess to shipping
var OrderShipped = 2

// OrderCompleted order are finished
var OrderCompleted = 3
