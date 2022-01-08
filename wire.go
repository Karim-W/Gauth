//go:build wireinject
// +build wireinject

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gauth.com/dbmanager"
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

func LoadDBVarsFromEnv() *dbmanager.DbSettingsConfig {
	fmt.Println("portNumber")
	portNumber, err := strconv.Atoi(LoadFromEnv("DBPort"))
	if err != nil {
		log.Fatalln(err)
	}
	return &dbmanager.DbSettingsConfig{
		Server:       LoadFromEnv("DBServer"),
		Port:         portNumber,
		User:         LoadFromEnv("DBUser"),
		Password:     LoadFromEnv("DBPassword"),
		DatabaseName: LoadFromEnv("DBName"),
	}
}

func InitializeDB() dbmanager.DBContext {
	wire.Build(NewDB)
	return dbmanager.DBContext{}
}

func NewDB() dbmanager.DBContext {
	conf := LoadDBVarsFromEnv()
	d := dbmanager.NewDBConnection(*conf)
	return dbmanager.DBContext{Ctx: d}
}
