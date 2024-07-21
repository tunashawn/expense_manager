package repositories

import (
	"Expense_Manager/pkg/wallet_service/models"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
)

const (
	wallets = "wallets"
)

type MySQLWalletRepo interface {
	CreateWallet(wallet models.Wallet) error
	UpdateWallet(wallet models.Wallet) error
	DeleteWallet(wallet models.Wallet) error
}

type MySQLWalletRepoImpl struct {
	db *bun.DB
}

func (w *MySQLWalletRepoImpl) DeleteWallet(wallet models.Wallet) error {
	_, err := w.db.NewDelete().
		Model(&wallet).
		Where("id = ?", wallet.ID).
		Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("could not delete wallet id=%d", wallet.ID))
	}

	return nil
}

func (w *MySQLWalletRepoImpl) UpdateWallet(wallet models.Wallet) error {
	_, err := w.db.NewUpdate().
		Model(&wallet).
		WherePK().
		Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("could not update wallet id=%d", wallet.ID))
	}

	return nil
}

func (w *MySQLWalletRepoImpl) CreateWallet(wallet models.Wallet) error {
	_, err := w.db.NewInsert().
		Model(&wallet).
		Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, "could not insert new wallet")
	}

	return nil
}
