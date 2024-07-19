package config

import (
	"io"
	"os"
	"time"

	"github.com/charmbracelet/log"
)

type Logger = log.Logger

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.NewWithOptions(writer, log.Options{
		Prefix:          p,
		Level:           log.DebugLevel,
		TimeFormat:      time.Kitchen,
		ReportTimestamp: true,
	})
	return logger
}
