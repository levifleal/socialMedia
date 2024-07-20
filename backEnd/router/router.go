package router

import (
	"github.com/gin-gonic/gin"
	"github.com/levifleal/socialMedia/backEnd/config"
)

var (
	logger *config.Logger
)

func Init() {
	logger = config.GetLogger("[Router]")
	logger.Debug("initializing Router...")

	r := gin.Default()

	initRoutes(r)

	r.Run()
}
