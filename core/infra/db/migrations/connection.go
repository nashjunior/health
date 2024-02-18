package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {

	fmt.Println("Starting connection")
	m, err := migrate.New(
		"file:///home/go/app/core/infra/db/migrations",
		"postgres://postgres:pmce@192.168.64.1:5432/health?sslmode=disable")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Starting migrations")
	err = m.Up()

	if err != nil {
		fmt.Println(err)

		fmt.Println("Rolling back migrations")
		err := m.Down()

		panic(err)

	}

	os.Exit(0)
}
