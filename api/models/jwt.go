package models

// JWT struct
type JWT struct {
	Token      string `json:"token"`
	Expiration int64  `json:"expiration"`
}
