package models

type Wallet struct {
	Name         string        `json:"name"`
	Balance      float64       `json:"balance"`
	Currency     string        `json:"currency"`
	Transactions []Transaction `json:"transactions,omitempty"`
}
