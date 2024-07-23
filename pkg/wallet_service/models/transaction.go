package models

import (
	"github.com/pkg/errors"
	"time"
)

const (
	income  = "Income"
	expense = "Expense"
)

type Transaction struct {
	ID         int       `json:"id" bun:"id,pk,autoincrement"`
	Type       string    `json:"type"`
	WalletID   int       `json:"wallet_id"`
	CategoryID int       `json:"category_id"`
	Amount     float64   `json:"amount"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
}

func (t *Transaction) VerifyTransaction() error {
	if t.ID < 1 {
		return errors.New("invalid ID")
	}

	if t.Type != income && t.Type != expense {
		return errors.New("invalid type")
	}

	if t.WalletID < 1 {
		return errors.New("invalid wallet_id")
	}

	if t.CategoryID < 1 {
		return errors.New("invalid category_id")
	}

	if t.Amount <= 0 {
		return errors.New("invalid wallet_id")
	}

	if t.Timestamp == (time.Time{}) {
		return errors.New("invalid timestamp")
	}

	return nil
}
