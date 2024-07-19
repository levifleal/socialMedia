package config

var (
	logger *Logger
)

func Init() error {

	logger.Debug("initializing configs...")


	return nil
}

func GetLogger(p string) *Logger {
	// initializer Logger
	logger = NewLogger(p)
	return logger
}
