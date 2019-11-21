package utils

// IDQuery helper that might be often used
type IDQuery struct {
	ID ID `json:"id"`
}

// ID type for id
type ID int64

type AccessTokenData struct {
	Token string `json:"token"`
}
