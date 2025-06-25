package db

import (
    "database/sql"
    _ "github.com/lib/pq"
)

var PG *sql.DB

func InitPostgres() {
    var err error
    PG, err = sql.Open("postgres", "postgres://user:pass@postgres:5432/urlshortener?sslmode=disable")
    if err != nil {
        panic(err)
    }
}