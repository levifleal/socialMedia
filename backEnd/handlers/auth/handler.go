package auth

import (
	"github.com/levifleal/socialMedia/backEnd/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitAuthHandler() {
	logger = config.GetLogger("[AuthHandler]")
	db = config.GetDB()
}
