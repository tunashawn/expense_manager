package db

import (
	"Expense_Manager/commons/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func NewMySQLConnection() (*bun.DB, error) {
	cfg := new(config.MySQLConfig)
	err := config.GetConfig(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "could not get mysql config")
	}

	sqldb, err := sql.Open("mysql", cfg.ConnectionString)
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, mysqldialect.New())

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "could not ping to mysql database")
	}

	return db, nil
}
