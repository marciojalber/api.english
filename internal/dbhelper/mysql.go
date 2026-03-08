// internal/dbhelper/mysql.go

package dbhelper

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/marciojalber/api.english/internal/config"
)

func MyCon() (*sql.DB, error) {
	db := config.Load().DB
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

func Select(sql string, cols []string) string {
	return strings.Replace(sql, ":cols", strings.Join(cols, ", "), 1)
}
