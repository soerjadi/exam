package models

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

// Category model
type Category struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	ParentID null.Int  `json:"parent_id"`
	Created  time.Time `json:"created"`
	Updated  null.Time `json:"updated"`
}
