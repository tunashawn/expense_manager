package repositories

import (
	db2 "Expense_Manager/commons/db"
	"Expense_Manager/pkg/user_manager_service/models"
	"context"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"log/slog"
)

type UserManagerMySQLRepo interface {
	CreateNewUser(user models.User) error
	IsUsernameExist(username string) (bool, error)
	IsEmailExist(email string) (bool, error)
	UpdateUserPassword(username string, password []byte) error
}

type UserManagerMySQLRepoImpl struct {
	db *bun.DB
}

func (u *UserManagerMySQLRepoImpl) CreateNewUser(user models.User) error {
	res, err := u.db.NewInsert().
		Model(&user).
		Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, "could not perform insert new user with username="+user.Username)
	}

	numRows, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "could not get num of rows affected")
	}

	if numRows == 1 {
		slog.Info("inserted a new user", "username", user.Username)
	}

	return nil
}

func (u *UserManagerMySQLRepoImpl) IsUsernameExist(username string) (bool, error) {
	count, err := u.db.NewSelect().
		Table("users").
		Where("username = ?", username).
		Count(context.Background())
	if err != nil {
		return false, errors.Wrap(err, "could not perform select")
	}

	if count == 1 {
		return true, nil
	}

	return false, nil
}

func (u *UserManagerMySQLRepoImpl) IsEmailExist(email string) (bool, error) {
	count, err := u.db.NewSelect().
		Table("users").
		Where("email = ?", email).
		Count(context.Background())
	if err != nil {
		return false, errors.Wrap(err, "could not perform select")
	}

	if count == 1 {
		return true, nil
	}

	return false, nil
}

func (u *UserManagerMySQLRepoImpl) UpdateUserPassword(username string, password []byte) error {
	res, err := u.db.NewInsert().
		Table("users").
		Set("password = ?", string(password)).
		Where("username = ?", username).
		Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, "could not perform update user password with username="+username)
	}

	numRows, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "could not get num of rows affected")
	}

	slog.Info("updated password of a user", "username", username, "rows_affected", numRows)

	return nil
}

func NewUserManagerMySQLRepo() (UserManagerMySQLRepo, error) {
	db, err := db2.NewMySQLConnection()
	if err != nil {
		return nil, err
	}

	return &UserManagerMySQLRepoImpl{db: db}, nil
}
