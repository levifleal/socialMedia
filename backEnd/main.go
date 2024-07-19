package main

import (
	"github.com/levifleal/socialMedia/backEnd/config"
	"github.com/levifleal/socialMedia/backEnd/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("[MAIN]")
	logger.Debug("initializating Api...")

	//initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	router.Init()
}
