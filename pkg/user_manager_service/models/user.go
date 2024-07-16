package models

type User struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	TotalBalance float64  `json:"total_balance"`
	Wallets      []Wallet `json:"wallets,omitempty"`
}

func (u *User) CalculateTotalBalance() {
	for _, wallet := range u.Wallets {
		u.TotalBalance += wallet.Balance
	}
}
