package logging

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger() {
	Logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func Info(message string) {
	Logger.Println("[INFO] " + message)
}

func Error(message string) {
	Logger.Println("[ERROR] " + message)
}

func Debug(message string) {
	Logger.Println("[DEBUG] " + message)
}
