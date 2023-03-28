package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

var db *sql.DB

func main() {
	dtbs, err := sql.Open("mysql", "root:root@/test")
	if err != nil {
		panic(err.Error()) // just for example
	}
	db = dtbs

	http.HandleFunc("/get_something", myHandler)

	println("listening..")
	http.ListenAndServe(":5005", nil)
}

func myHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := db.QueryContext(ctx, "SELECT sleep(15)")
	if err != nil {
		log.Println("error query: " + err.Error())
		w.Write([]byte("something's wrong: " + err.Error()))
		return
	}

	w.Write([]byte("success"))
}
