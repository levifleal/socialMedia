package config

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	logger *Logger
	db     *gorm.DB
)

func Init() error {
	var err error
	logger.Debug("initializing configs...")

	//initialize dotEnv
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//initializer mySql
	db, err = initMySql()
	if err != nil {
		return fmt.Errorf("error initializing mySql: %v", err.Error())
	}

	return nil
}

func GetLogger(p string) *Logger {
	// initializer Logger
	logger = NewLogger(p)
	return logger
}

func GetDB() *gorm.DB {
	return db
}
