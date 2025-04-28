package db

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
    var err error
    DB, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/curso_go_db")
    if err != nil {
        log.Fatal("Failed to open database:", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    log.Println("Database connected!")
}
