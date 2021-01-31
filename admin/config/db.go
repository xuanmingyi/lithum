package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitDB() (err error) {
	Config.DB, err = sql.Open("mysql", Config.DSN)
	if err != nil {
		panic(err)
	}
	return nil
}
