package config

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initMySql() (*gorm.DB, error) {
	logger := GetLogger("[MySQL]")

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	ip := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, pass, ip, dbName)

	if user == "" || pass == "" || ip == "" || dbName == "" {
		logger.Errorf("don't have all the needed vallues in dotEnv")
		return nil, fmt.Errorf("don't have all the needed vallues")
	}

	logger.Debug("initializing database connection...")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("cannot connect to mySQL: %s", err.Error())
		return nil, err
	}

	return db, nil
}
