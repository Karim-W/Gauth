package dbmanager

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewDBConnection(conf DbSettingsConfig) *gorm.DB {
	return StartConnection(conf)
}

func StartConnection(config DbSettingsConfig) *gorm.DB {
	// Build connection string
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		config.User, config.Password, config.Server, config.Port, config.DatabaseName)
	var err error
	fmt.Println("Connecting to database...")
	fmt.Println(connString)
	// Create connection pool
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

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

type DBContext struct {
	Ctx *gorm.DB
}
