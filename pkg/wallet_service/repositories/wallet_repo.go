package repositories

import (
	db2 "Expense_Manager/commons/db"
)

type WalletRepository struct {
	Wallet      MySQLWalletRepo
	transaction MySQLTransactionRepo
	report      MySQLReportRepo
}

func NewRepository() (*WalletRepository, error) {
	db, err := db2.NewMySQLConnection()
	if err != nil {
		return nil, err
	}

	return &WalletRepository{
		Wallet:      &MySQLWalletRepoImpl{db: db},
		transaction: &MySQLTransactionRepoImpl{db: db},
		report:      &MySQLReportRepoImpl{db: db},
	}, nil
}
