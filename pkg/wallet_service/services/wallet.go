package services

import (
	"Expense_Manager/pkg/wallet_service/models"
	"Expense_Manager/pkg/wallet_service/repositories"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
)

type WalletService interface {
	CreateWallet(wallet models.Wallet) error
	UpdateWallet(wallet models.Wallet) error
	DeleteWallet(id int) error
	GetWalletList(userID int) ([]models.Wallet, error)
	GetWallet(walletID int) (*models.Wallet, error)
}

type WalletServiceImpl struct {
	sqlWalletRepo repositories.SqlWalletRepo
}

func (w *WalletServiceImpl) CreateWallet(wallet models.Wallet) error {
	return w.sqlWalletRepo.CreateWallet(wallet)
}

func (w *WalletServiceImpl) UpdateWallet(wallet models.Wallet) error {
	return w.sqlWalletRepo.CreateWallet(wallet)
}

func (w *WalletServiceImpl) DeleteWallet(id int) error {
	return w.sqlWalletRepo.DeleteWallet(id)
}

func (w *WalletServiceImpl) GetWalletList(userID int) ([]models.Wallet, error) {
	//TODO implement me
	panic("implement me")
}

func (w *WalletServiceImpl) GetWallet(walletID int) (*models.Wallet, error) {
	return w.sqlWalletRepo.GetWallet(walletID)
}

func NewWalletService(db *bun.DB) (WalletService, error) {
	sqlWalletRepo, err := repositories.NewSqlWalletRepo(db)
	if err != nil {
		return nil, errors.Wrap(err, "could not init sql wallet repo")
	}
	return &WalletServiceImpl{sqlWalletRepo: sqlWalletRepo}, nil
}
