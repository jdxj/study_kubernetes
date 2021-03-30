package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	dbIns *sqlx.DB
)

func Init(user, pass, host, port, dbName string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=Local&parseTime=true",
		user, pass, host, port, dbName)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}

	dbIns, err = db, db.Ping()
	return err
}
