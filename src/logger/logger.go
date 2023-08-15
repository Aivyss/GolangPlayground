package logger

import (
	"log"
	"os"
)

func GetLogger() (*log.Logger, *os.File) {
	logFile, err := os.OpenFile(
		"log.log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}

	logger := log.New(logFile, "Custom Logger: ", log.Ldate|log.Ltime|log.Lshortfile)

	return logger, logFile
}
