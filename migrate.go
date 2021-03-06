package main

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func pgConnectionString() string {
	user := os.Getenv("PG_USER")
	pass := os.Getenv("PG_PASS")
	dbas := os.Getenv("PG_DBAS")
	port := os.Getenv("PG_PORT")
	host := os.Getenv("PG_HOST")
	ssl := os.Getenv("PG_SSL")
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, pass, host, port, dbas, ssl)
}

func migrateUp() {
	m, err := migrate.New("file://migrations", pgConnectionString())
	if err != nil {
		panic(err)
	}
	// migration upto 10 steps
	m.Steps(10)
}
