package main

import (
	"errors"
	"log"
	"vinbigdata/config"
	"vinbigdata/database/migration"
	"vinbigdata/internal/delivery/http"
	"vinbigdata/package/db"
)

func main() {

	// connect to database
	database, err := db.NewDatabase(config.DbConnStr, 10, 2)
	if err != nil {
		log.Fatal(errors.New("no db connection"))
	}
	defer db.CloseDatabase(database)
	migration.CreateTable(database)

	app := http.InitHttp(database)
	app.Run(":80")
}
