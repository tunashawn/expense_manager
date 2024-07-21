package repositories

import (
	"Expense_Manager/pkg/wallet_service/models"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"time"
)

type MySQLTransactionRepo interface {
	AddTransaction(transaction models.Transaction) error
	UpdateTransaction(transaction models.Transaction) error
	DeleteTransaction(transaction models.Transaction) error
	GetTransactionsByMonth(walletID int, month int) ([]models.Transaction, error)
	GetTransactionsByDate(walletID int, date time.Time) ([]models.Transaction, error)
}

type MySQLTransactionRepoImpl struct {
	db *bun.DB
}

func (t *MySQLTransactionRepoImpl) GetTransactionsByDate(walletID int, date time.Time) ([]models.Transaction, error) {
	var res []models.Transaction
	_, err := t.db.NewSelect().
		Model(&res).
		Where("wallet_id = ? AND DATE(timestamp) = DATE(?)", walletID, date).
		Order("timestamp DESC").
		Exec(context.Background(), &res)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("could not get transactions for wallet_id=%d and date=%s", walletID, date.Format(time.DateOnly)))
	}

	return res, nil
}

func (t *MySQLTransactionRepoImpl) GetTransactionsByMonth(walletID int, month int) ([]models.Transaction, error) {
	var res []models.Transaction
	_, err := t.db.NewSelect().
		Model(&res).
		Where("wallet_id = ? AND MONTH(timestamp) = ?", walletID, month).
		Order("timestamp DESC").
		Exec(context.Background(), &res)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("could not get transactions for wallet_id=%d and month=%d", walletID, month))
	}

	return res, nil
}

func (t *MySQLTransactionRepoImpl) AddTransaction(transaction models.Transaction) error {
	_, err := t.db.NewInsert().
		Model(&transaction).
		Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, "could not insert new transaction")
	}

	return nil
}

func (t *MySQLTransactionRepoImpl) UpdateTransaction(transaction models.Transaction) error {
	_, err := t.db.NewUpdate().
		Model(&transaction).
		WherePK().
		Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, "could not update new transaction")
	}

	return nil
}

func (t *MySQLTransactionRepoImpl) DeleteTransaction(transaction models.Transaction) error {
	_, err := t.db.NewDelete().
		Model(&transaction).
		WherePK().
		Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, "could not delete new transaction")
	}

	return nil
}
