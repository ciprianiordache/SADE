package models

type Transaction struct {
	Token       string  `json:"token"`
	AdminID     int64   `json:"admin_id"`
	MediaID     int64   `json:"media_id"`
	ClientEmail string  `json:"client_email"`
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Description string  `json:"description"`
}
