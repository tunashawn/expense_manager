package models

type Transaction struct {
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}
