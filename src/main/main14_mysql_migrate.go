package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
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
		if err.Error() != "no change" {
			panic(err)
		}
	}

	db2, _ := sql.Open("mysql", "root:sany_root@tcp(10.162.4.9:3306)/sany_scada?multiStatements=true")
	driver2, _ := mysql.WithInstance(db2, &mysql.Config{})
	m2, err2 := migrate.NewWithDatabaseInstance(
		"file://migrations/sany_scada/",
		"mysql",
		driver2,
	)
	if err2 != nil {
		panic(err2)
	}
	err2 = m2.Up()
	if err2 != nil {
		panic(err2)
	}

}
