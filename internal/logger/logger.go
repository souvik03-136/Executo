package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger
)

// Init initializes the logrus logger.
func Init(logFilePath string) error {
	// Create or open the log file
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// Create a new logrus logger
	log = logrus.New()

	// Set output for the logger
	log.SetOutput(file)

	// Set log format (can use JSON or Text format)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return nil
}

// Info logs informational messages.
func Info(message string) {
	if log != nil {
		log.Info(message)
	}
}

// Error logs error messages.
func Error(message string) {
	if log != nil {
		log.Error(message)
	}
}
