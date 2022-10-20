package config

import (
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

var lock = &sync.Mutex{}

type AppConfig struct {
	DBPort    uint
	DBUser    string
	DBPwd     string
	DBHost    string
	DBName    string
	JWTSecret string
}

func NewConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()
	cfg := initConfig()
	if cfg == nil {
		log.Fatal("Cannot run configuration setup")
		return nil
	}
	return cfg
}

func initConfig() *AppConfig {
	var app AppConfig

	/* Buka tag comment untuk run app di localhost. */
	godotenv.Load("config.env")

	app.DBUser = os.Getenv("DB_USER")
	app.DBPwd = os.Getenv("DB_PWD")
	app.DBHost = os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Error("config error :", err.Error())
		return nil
	}
	app.DBPort = uint(port)
	app.DBName = os.Getenv("DB_NAME")
	app.JWTSecret = os.Getenv("JWT_SECRET")

	return &app
}
