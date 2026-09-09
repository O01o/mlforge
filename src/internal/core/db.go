package core

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
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
	tlsEnabled := false
	if os.Getenv("DB_TLS") != "" {
		var err error
		tlsEnabled, err = strconv.ParseBool(os.Getenv("DB_TLS"))
		if err != nil {
			log.Fatal(err)
		}
	}

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
	if config.TLS {
		if config.CACert != "" {
			rootCertPool := x509.NewCertPool()
			caCert, err := os.ReadFile(config.CACert)
			if err != nil {
				log.Fatal(err)
			}
			if ok := rootCertPool.AppendCertsFromPEM(caCert); !ok {
				log.Fatal("Failed to append PEM.")
			}
			tlsConfig := &tls.Config{
				RootCAs: rootCertPool,
			}
			// Setting mTLS
			if config.ClientCert != "" && config.ClientKey != "" {
				clientCert, err := tls.LoadX509KeyPair(config.ClientCert, config.ClientKey)
				if err != nil {
					log.Fatal(err)
				}
				tlsConfig.Certificates = []tls.Certificate{clientCert}
			}
			mysql.RegisterTLSConfig("custom", tlsConfig)
			dsn += "&tls=custom"
		} else {
			dsn += "&tls=true"
		}
	}
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
