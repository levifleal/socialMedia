package config

import (
	"fmt"
	"os"

	"github.com/levifleal/socialMedia/backEnd/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initMySql() (*gorm.DB, error) {
	logger := GetLogger("[MySQL]")

	//defining fields
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	ip := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	
	//verify if has every field
	if user == "" || pass == "" || ip == "" || dbName == "" {
		logger.Errorf("don't have all the needed vallues in dotEnv")
		return nil, fmt.Errorf("don't have all the needed vallues")
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", user, pass, ip, dbName)

	logger.Debug("initializing database connection...")

	//connecting with mySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("cannot connect to mySQL: %s", err.Error())
		return nil, err
	}

	//Migrate the User Schema
	err = db.AutoMigrate(&schemas.User{})
	if err != nil {
		logger.Errorf("sqlite autoMigration error: %v", err)
	}

	return db, nil
}
