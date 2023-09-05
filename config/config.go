package config

var (
	logger *Logger
)

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}