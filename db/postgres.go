package db

import (
    "database/sql"
    _ "github.com/lib/pq"
)

var PG *sql.DB

/*
 * @brief InitPostgres initializes a PostgreSQL connection pool.
 *
 * This function opens a PostgreSQL connection pool using the environment variables
 * PG_HOST, PG_USER, PG_PASSWORD, PG_DBNAME, and PG_PORT. The pool is stored in the
 * package variable PG.
 */
func InitPostgres() {
    var err error
    PG, err = sql.Open("postgres", "postgres://user:pass@postgres:5432/urlshortener?sslmode=disable")
    if err != nil {
        panic(err)
    }
}