package logger

import (
	"log"
	"os"
)

var (
	Trace   *log.Logger // Just about anything
	Info    *log.Logger // Important information
	Warning *log.Logger // Be concerned
	Error   *log.Logger // Critical problem
)

func init() {
	Trace = log.New(os.Stderr, "TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stderr, "INFO:  ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stderr, "WARN:  ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(os.Stderr, "ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
