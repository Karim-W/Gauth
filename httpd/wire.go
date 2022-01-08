package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	dbmanager "gauth.com/httpd/DBManager"
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
