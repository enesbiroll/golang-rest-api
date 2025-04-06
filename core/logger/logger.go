package logger

import (
	"fmt"
	"io"
	"os"
	"rest-api/config"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Global logger
var Log *logrus.Logger

// LogHook for logging to the database
type LogHook struct {
	DB *gorm.DB
}

// LogHook levels for warning and fatal only
func (hook *LogHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.WarnLevel,
		logrus.FatalLevel,
	}
}

// Log writes to the database
func (hook *LogHook) Fire(entry *logrus.Entry) error {
	// Only log 'warn' and 'fatal' levels
	if entry.Level == logrus.WarnLevel || entry.Level == logrus.FatalLevel {
		logRecord := struct {
			Level   string `json:"level"`
			Message string `json:"message"`
		}{
			Level:   entry.Level.String(),
			Message: entry.Message,
		}

		// Insert into DB
		if err := hook.DB.Create(&logRecord).Error; err != nil {
			return fmt.Errorf("failed to insert log into database: %v", err)
		}
	}

	return nil
}

// Initialize the logger
func Init() {
	// Check if Log is already initialized
	if Log != nil {
		return
	}

	// Initialize a new logrus instance
	Log = logrus.New()

	// Set log level to info
	Log.SetLevel(logrus.InfoLevel)

	// Open log file
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		Log.Error("Could not open log file, using console logging instead")
	} else {
		Log.SetOutput(io.MultiWriter(os.Stdout, file)) // Write to both stdout and file
	}

	// Set log format
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Add DB hook for warning and fatal logs
	Log.AddHook(&LogHook{DB: config.DB})
}
