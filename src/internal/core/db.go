package core

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

func ConnectDB() *sqlx.DB {
	dsn := os.Getenv("DATA_SOURCE_NAME")
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db
}
