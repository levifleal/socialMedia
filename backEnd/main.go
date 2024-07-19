package main

import "github.com/levifleal/socialMedia/backEnd/config"

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("[MAIN]")

	//initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	logger.Debug("initializating Api...")
}
