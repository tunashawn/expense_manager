package repository

import (
	db2 "Expense_Manager/commons/db"
	"context"
	"github.com/uptrace/bun"
)

const users = "users"

type AuthServiceMySQLRepo interface {
	GetHashedPasswordOfUser(username string) ([]byte, error)
}

type AuthServiceMySQLRepoImpl struct {
	db *bun.DB
}

func (a *AuthServiceMySQLRepoImpl) GetHashedPasswordOfUser(username string) ([]byte, error) {
	var password string
	count, err := a.db.NewSelect().
		Column("password").
		Table(users).
		Where("username = ?", username).
		ScanAndCount(context.Background(), &password)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, nil
	}

	return []byte(password), nil
}

func NewAuthServiceMySQLRepository() (AuthServiceMySQLRepo, error) {
	db, err := db2.NewMySQLConnection()
	if err != nil {
		return nil, err
	}

	return &AuthServiceMySQLRepoImpl{db: db}, nil
}
