// internal/dbhelper/mysql.go

package handler

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DBHandler struct{}

var DB DBHandler = DBHandler{}

func (DB *DBHandler) MyCon() (*sql.DB, error) {
	db := Config.Load().DB
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		db["default"].User,
		db["default"].Password,
		db["default"].Host,
		db["default"].Port,
		db["default"].DBName,
	)

	res, err := sql.Open(db["default"].Driver, dsn)
	if err != nil {
		return nil, err
	}

	err = res.Ping()
	if err != nil {
		return nil, err
	}

	return res, nil
}
