package core

import (
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type dbConfig struct {
	Host       string
	Port       string
	User       string
	Password   string
	Name       string
	TLS        bool
	CACert     string
	ClientCert string
	ClientKey  string
}

func ConnectDB() *sqlx.DB {
	tlsEnabled, _ := strconv.ParseBool(os.Getenv("DB_TLS"))

	config := &dbConfig{
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
		Name:       os.Getenv("DB_NAME"),
		TLS:        tlsEnabled,
		CACert:     os.Getenv("DB_CA_CERT"),
		ClientCert: os.Getenv("DB_CLIENT_CERT"),
		ClientKey:  os.Getenv("DB_CLIENT_KEY"),
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&multiStatements=true",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
