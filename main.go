package main

import (
	"log"
	"os"

	"github.com/juanrod1028/Transactions/src/adapter/postgres"
	"github.com/juanrod1028/Transactions/src/usecase/service"
)

func main() {
	os.Setenv("COMPANY_EMAIL", "")
	os.Setenv("COMPANY_EMAIL_PASS", "")
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "password"
		dbname   = "postgres"
	)
	postgresStorage, err := postgres.NewPostgresStore(host, port, user, password, dbname)
	if err != nil {
		log.Fatal(err)
	}
	if err := postgresStorage.Init(); err != nil {
		log.Fatal(err)
	}

	server := service.NewApiServer(":3000", postgresStorage)
	server.Run()
}
