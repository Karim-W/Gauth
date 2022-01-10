//go:build wireinject
// +build wireinject

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gauth.com/Database"
	"github.com/google/wire"
	"github.com/joho/godotenv"
)

func LoadFromEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)

}

func LoadDBVarsFromEnv() *Database.DbSettingsConfig {
	fmt.Println("portNumber")
	portNumber, err := strconv.Atoi(LoadFromEnv("DBPort"))
	if err != nil {
		log.Fatalln(err)
	}
	return &Database.DbSettingsConfig{
		Server:       LoadFromEnv("DBServer"),
		Port:         portNumber,
		User:         LoadFromEnv("DBUser"),
		Password:     LoadFromEnv("DBPassword"),
		DatabaseName: LoadFromEnv("DBName"),
	}
}

func InitializeDB() Database.DBContext {
	wire.Build(NewDB)
	return Database.DBContext{}
}

func NewDB() Database.DBContext {
	conf := LoadDBVarsFromEnv()
	d := Database.NewDBConnection(*conf)
	return Database.DBContext{Ctx: d}
}
