package dbmanager

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var db *sql.DB

func newDBConnection(conf DbSettingsConfig) *sql.DB {
	return StartConnection(conf)
}

func StartConnection(config DbSettingsConfig) *sql.DB {
	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		config.Server, config.User, config.Password, config.Port, config.DatabaseName)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("Connected!")
	return db
}

type DbSettingsConfig struct {
	Server       string
	Port         int
	User         string
	Password     string
	DatabaseName string
}
