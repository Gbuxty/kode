package utils

import (
	"encoding/json"
	
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	// Настройка logrus
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
}

// LogInfo logs informational messages
func LogInfo(message string, fields map[string]interface{}) {
	logMessage("INFO", message, fields)
}

// LogError logs error messages
func LogError(message string, fields map[string]interface{}) {
	logMessage("ERROR", message, fields)
}

// logMessage logs messages in a unified format
func logMessage(level, message string, fields map[string]interface{}) {
	logEntry := map[string]interface{}{
		"level":   level,
		"message": message,
		"time":    time.Now().Format(time.RFC3339),
	}
	for k, v := range fields {
		logEntry[k] = v
	}
	logJSON, err := json.Marshal(logEntry)
	if err != nil {
		log.Printf("ERROR: failed to marshal log entry: %v", err)
		return
	}
	log.Println(string(logJSON))
}