package logger

import (
	"log"
	"os"
)

// LogErrorToFile appends the error message to the "error.log" file.
func LogErrorToFile(input string, err error) {
	// Open the file in append mode. If the file doesn't exist, create it.
	file, err := os.OpenFile("../logs/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error opening the file:", err)
	}
	defer file.Close()

	// Create a new logger with the file as the output destination.
	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	// Write the error message to the log file.
	logger.Println(input, err)
}

// LogAccessToFile appends the access message to the "error.log" file.
func LogAccessToFile(access string) {
	// Open the file in append mode. If the file doesn't exist, create it.
	file, err := os.OpenFile("../logs/access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal("Error opening the file:", err)
	}
	defer file.Close()

	// Create a new logger with the file as the output destination.
	logger := log.New(file, "", log.Ldate|log.Ltime|log.Lmicroseconds)

	// Write the access message to the log file.
	logger.Println(access)
}