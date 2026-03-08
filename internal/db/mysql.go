// internal/db/mysql.go

package db

import (
    "fmt"
    "database/sql"
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

	return sql.Open(db["default"].Driver, dsn)
}

/*
--- QUERY ---

rows, err := db.Query("SELECT id, name FROM users")
if err != nil {
    panic(err)
}
defer rows.Close()

for rows.Next() {
    var id int
    var name string

    rows.Scan(&id, &name)

    fmt.Println(id, name)
}



--- INSERT ---

result, err := db.Exec(
    "INSERT INTO users(name) VALUES(?)",
    "Alice",
)

if err != nil {
    panic(err)
}

id, _ := result.LastInsertId()
fmt.Println("Inserted ID:", id)
*/
