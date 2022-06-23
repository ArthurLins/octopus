package logger

import "log"

func Error(message string) {
	log.Println("[ERROR] - " + message)
}

func Info(message string) {
	log.Println("[INFO] - " + message)
}
