package main

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, _ := sql.Open("mysql", "root:sany_root@tcp(10.162.4.9:3306)/sany_event_center?multiStatements=true")
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"mysql",
		driver,
	)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil {
		panic(err)
	}

}
