package logs

import (
	"log"
	"os"
)

func Setlogger() {
	LoggerFile, err := os.OpenFile("Loginmodule Logger.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("Failed to create file")
	}
	log.SetOutput(LoggerFile)
}
