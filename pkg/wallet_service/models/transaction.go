package models

import "time"

type Transaction struct {
	ID         int       `json:"id" bun:"id,pk,autoincrement"`
	Type       string    `json:"type"`
	WalletID   int       `json:"wallet_id"`
	CategoryID int       `json:"category_id"`
	Amount     float64   `json:"amount"`
	Timestamp  time.Time `json:"timestamp"`
}
