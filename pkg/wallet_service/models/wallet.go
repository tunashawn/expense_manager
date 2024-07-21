package models

import "time"

type Wallet struct {
	ID           int           `json:"id" bun:"id,pk,autoincrement"`
	Name         string        `json:"name"`
	Balance      float64       `json:"balance"`
	Currency     int           `json:"currency"`
	Transactions []Transaction `json:"transactions,omitempty" bun:"-"`
	CreatedAt    time.Time     `json:"created_at"`
	IconID       int           `json:"icon_id"`
}
