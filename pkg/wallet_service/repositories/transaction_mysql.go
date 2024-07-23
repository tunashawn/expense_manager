package repositories

import (
	"Expense_Manager/pkg/wallet_service/models"
	"context"
	"database/sql"
	errors2 "errors"
	"fmt"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"log/slog"
	"time"
)

type SqlTransactionRepo interface {
	AddTransaction(transaction models.Transaction) error
	UpdateTransaction(transaction models.Transaction) error
	DeleteTransaction(transaction models.Transaction) error
	GetTransactionListByMonth(walletID int, month int) ([]models.Transaction, error)
	GetTransactionListByDate(walletID int, date time.Time) ([]models.Transaction, error)
	GetTransaction(transactionID int) (*models.Transaction, error)
}

type SqlTransactionRepoImpl struct {
	db *bun.DB
}

func (t *SqlTransactionRepoImpl) GetTransaction(transactionID int) (*models.Transaction, error) {
	var res models.Transaction
	_, err := t.db.NewSelect().
		Model(&res).
		Where("id = ?", transactionID).
		Exec(context.Background(), &res)
	if err != nil && !errors2.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrap(err, fmt.Sprintf("could not get transactions for id=%d", transactionID))
	}

	return &res, nil
}

func (t *SqlTransactionRepoImpl) GetTransactionListByDate(walletID int, date time.Time) ([]models.Transaction, error) {
	var res []models.Transaction
	_, err := t.db.NewSelect().
		Model(&res).
		Where("wallet_id = ? AND DATE(timestamp) = DATE(?)", walletID, date).
		Order("timestamp DESC").
		Exec(context.Background(), &res)
	if err != nil && !errors2.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrap(err, fmt.Sprintf("could not get transactions for wallet_id=%d and date=%s", walletID, date.Format(time.DateOnly)))
	}

	return res, nil
}

func (t *SqlTransactionRepoImpl) GetTransactionListByMonth(walletID int, month int) ([]models.Transaction, error) {
	var res []models.Transaction
	_, err := t.db.NewSelect().
		Model(&res).
		Where("wallet_id = ? AND MONTH(timestamp) = ?", walletID, month).
		Order("timestamp DESC").
		ScanAndCount(context.Background(), &res)
	if err != nil && !errors2.Is(err, sql.ErrNoRows) {
		return nil, errors.Wrap(err, fmt.Sprintf("could not get transactions for wallet_id=%d and month=%d", walletID, month))
	}

	return res, nil
}

func (t *SqlTransactionRepoImpl) AddTransaction(transaction models.Transaction) error {
	result, err := t.db.NewInsert().
		Model(&transaction).
		Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, "could not insert new transaction")
	}

	numOfRowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "could not get num of rows affected")
	}

	slog.Info("inserted a new transaction", "numOfRowsAffected", numOfRowsAffected)

	return nil
}

func (t *SqlTransactionRepoImpl) UpdateTransaction(transaction models.Transaction) error {
	result, err := t.db.NewUpdate().
		Model(&transaction).
		WherePK().
		Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, "could not update new transaction")
	}

	numOfRowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "could not get num of rows affected")
	}

	slog.Info("updated a transaction", "numOfRowsAffected", numOfRowsAffected)

	return nil
}

func (t *SqlTransactionRepoImpl) DeleteTransaction(transaction models.Transaction) error {
	result, err := t.db.NewDelete().
		Model(&transaction).
		WherePK().
		Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, "could not delete new transaction")
	}

	numOfRowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "could not get num of rows affected")
	}

	slog.Info("deleted a transaction", "numOfRowsAffected", numOfRowsAffected)

	return nil
}

func NewSqlTransactionRepo(db *bun.DB) (SqlTransactionRepo, error) {
	return &SqlTransactionRepoImpl{db: db}, nil
}
