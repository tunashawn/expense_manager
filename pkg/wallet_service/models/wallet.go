package models

import (
	"github.com/pkg/errors"
	"time"
)

type Wallet struct {
	ID           int           `json:"id" bun:"id,pk,autoincrement"`
	Name         string        `json:"name"`
	Balance      float64       `json:"balance"`
	CurrencyID   int           `json:"currency_id"`
	Transactions []Transaction `json:"transactions,omitempty" bun:"-"`
	CreatedAt    time.Time     `json:"created_at"`
	IconID       int           `json:"icon_id"`
}

func (w *Wallet) VerifyWallet() error {
	if w.Name == "" {
		return errors.New("invalid wallet name")
	}

	if w.Balance < 0 {
		return errors.New("invalid wallet balance")
	}

	if w.CurrencyID <= 0 {
		return errors.New("invalid wallet currency")
	}

	if w.CreatedAt == (time.Time{}) {
		return errors.New("invalid wallet created_at")
	}

	if w.IconID <= 0 {
		return errors.New("invalid wallet icon")
	}

	return nil
}
