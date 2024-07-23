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
)

const (
	wallets = "wallets"
)

type SqlWalletRepo interface {
	CreateWallet(wallet models.Wallet) error
	UpdateWallet(wallet models.Wallet) error
	DeleteWallet(id int) error
	GetWalletList(userID int) ([]models.Wallet, error)
	GetWallet(walletID int) (*models.Wallet, error)
}

type SqlWalletRepoImpl struct {
	db *bun.DB
}

func (w *SqlWalletRepoImpl) GetWallet(walletID int) (*models.Wallet, error) {
	var wallet models.Wallet
	err := w.db.NewSelect().Model(&wallet).Where("id = ?", walletID).Scan(context.Background(), &wallet)
	if err != nil && !errors2.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return &wallet, nil
}

func (w *SqlWalletRepoImpl) GetWalletList(userID int) ([]models.Wallet, error) {
	var walletList []models.Wallet
	err := w.db.NewSelect().Model(&walletList).Where("user_id = ?", userID).Scan(context.Background(), &walletList)
	if err != nil && !errors2.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return walletList, nil
}

func (w *SqlWalletRepoImpl) DeleteWallet(id int) error {
	result, err := w.db.NewDelete().
		Table("wallets").
		Where("id = ?", id).
		Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("could not delete wallet id=%d", id))
	}

	numOfRowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "could not get num of rows affected")
	}

	slog.Info("deleted a new wallet", "numOfRowsAffected", numOfRowsAffected)

	return nil
}

func (w *SqlWalletRepoImpl) UpdateWallet(wallet models.Wallet) error {
	result, err := w.db.NewUpdate().
		Model(&wallet).
		WherePK().
		Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("could not update wallet id=%d", wallet.ID))
	}

	numOfRowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "could not get num of rows affected")
	}

	slog.Info("updated a new wallet", "numOfRowsAffected", numOfRowsAffected)

	return nil
}

func (w *SqlWalletRepoImpl) CreateWallet(wallet models.Wallet) error {
	result, err := w.db.NewInsert().
		Model(&wallet).
		Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, "could not insert new wallet")
	}

	numOfRowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "could not get num of rows affected")
	}

	slog.Info("added a new wallet", "numOfRowsAffected", numOfRowsAffected)

	return nil
}

func NewSqlWalletRepo(db *bun.DB) (SqlWalletRepo, error) {
	return &SqlWalletRepoImpl{db: db}, nil
}
