package logger

import (
	"io"
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func InitLogger() {

	// Create/Open log file
	file, err := os.OpenFile(
		"app.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)

	if err != nil {
		log.Fatal("Failed to create log file:", err)
	}

	// Write logs to both console and file
	multiWriter := io.MultiWriter(os.Stdout, file)

	InfoLogger = log.New(
		multiWriter,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	ErrorLogger = log.New(
		multiWriter,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
}
