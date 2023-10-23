package logger

import (
	"log"
	"os"
)

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Error(v ...interface{}) {
	l.SetPrefix("ERROR: ")
	l.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.SetPrefix("INFO: ")
	l.Println(v...)
}
