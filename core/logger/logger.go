package logger

import (
	"io"
	"os"

	Config "rest-api/config"
	Models "rest-api/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Initialize the logger
var Log *logrus.Logger

// LogHook for logging into the database for specific levels
type LogHook struct {
	DB *gorm.DB
}

func (hook *LogHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.WarnLevel,
		logrus.FatalLevel,
	}
}

func (hook *LogHook) Fire(entry *logrus.Entry) error {
	// Only log warning or fatal messages to the database
	if entry.Level == logrus.WarnLevel || entry.Level == logrus.FatalLevel {
		logRecord := Models.Log{
			Level:   entry.Level.String(),
			Message: entry.Message,
		}
		return hook.DB.Create(&logRecord).Error
	}
	return nil
}

// Initialize the logger
func Init() {
	// Create a new logger instance
	Log = logrus.New()

	// Set log level (you can change it to Debug, Info, Warn, etc.)
	Log.SetLevel(logrus.InfoLevel)

	// Create or open the log file
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		Log.Error("Could not open log file, using console logging instead")
	} else {
		// Log to both console and file using MultiWriter
		Log.SetOutput(io.MultiWriter(os.Stdout, file))
	}

	// Set log format (you can use JSON or Text format)
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Add the database logging hook for specific log levels (warn and fatal)
	Log.AddHook(&LogHook{DB: Config.DB})
}
