package services

import (
	"Expense_Manager/pkg/wallet_service/models"
	"Expense_Manager/pkg/wallet_service/repositories"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
)

type TransactionService interface {
	CreateNewTransaction(transaction models.Transaction) error
	GetTransaction(id int) (*models.Transaction, error)
	GetTransactionList() ([]models.Transaction, error)
	UpdateTransaction(transaction models.Transaction) error
	DeleteTransaction(id int) error
}

type TransactionServiceImpl struct {
	sqlTransactionRepo repositories.SqlTransactionRepo
}

func (t *TransactionServiceImpl) CreateNewTransaction(transaction models.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (t *TransactionServiceImpl) GetTransaction(id int) (*models.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TransactionServiceImpl) GetTransactionList() ([]models.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TransactionServiceImpl) UpdateTransaction(transaction models.Transaction) error {
	//TODO implement me
	panic("implement me")
}

func (t *TransactionServiceImpl) DeleteTransaction(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewTransactionService(db *bun.DB) (TransactionService, error) {
	sqlTransactionRepo, err := repositories.NewSqlTransactionRepo(db)
	if err != nil {
		return nil, errors.Wrap(err, "could not create sql transaction repo")
	}
	return &TransactionServiceImpl{sqlTransactionRepo: sqlTransactionRepo}, nil
}
