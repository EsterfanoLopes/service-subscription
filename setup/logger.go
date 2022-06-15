package setup

import (
	"log"
	"os"
)

func NewErrorLogger() *log.Logger {
	return log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}

func NewInfoLogger() *log.Logger {
	return log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
}
