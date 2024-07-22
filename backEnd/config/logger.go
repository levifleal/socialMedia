package config

import (
	"io"
	"os"
	"time"

	"github.com/charmbracelet/log"
)

type Logger = log.Logger

func NewLogger(p string) *Logger {
	var level log.Level
	writer := io.Writer(os.Stdout)

	//verify if is is prod mode
	if os.Getenv("MODE") == "prod" {
		level = log.InfoLevel
	} else {
		level = log.DebugLevel
	}

	//initializing logger
	logger := log.NewWithOptions(writer, log.Options{
		Prefix:          p,
		Level:           level,
		TimeFormat:      time.Kitchen,
		ReportTimestamp: true,
	})
	return logger
}
