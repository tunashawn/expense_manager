package repository

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"log"
	"log/slog"
	"testing"
)

func TestAuthServiceMySQLRepoImpl_GetHashedPasswordOfUser(t *testing.T) {
	sqldb, err := sql.Open("mysql", "root:@tcp(localhost:3306)/expense_manager")
	if err != nil {
		log.Fatal(err)
	}

	db := bun.NewDB(sqldb, mysqldialect.New())

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	x := AuthServiceMySQLRepoImpl{db: db}

	p, err := x.GetHashedPasswordOfUser("default")
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("password", "p", p)
}
