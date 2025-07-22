package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/JairoCC/go-react/backend/internal/repository"
	"github.com/JairoCC/go-react/backend/internal/repository/dbrepo"
)

const port = 8080

type application struct {
	DSN    string
	Domain string
	DB     repository.DatabaseRepo
}

func main() {
	// set application config
	var app application

	// read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5433 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")

	flag.Parse()

	// connect tro the database
	conn, err := app.connectioToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()
	// set the domain and port
	app.Domain = "example.com"
	log.Println("starting application on port: ", port)

	// start a web server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}
