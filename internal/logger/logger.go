package logger

import (
	"log"
	"os"
)
var Info *log.Logger
var Warn *log.Logger

func init() {
	Info = log.New(os.Stdout, "INFO:", log.Ltime | log.Lshortfile )
	Warn = log.New(os.Stdout, "!!!WARNING: ", log.Ldate | log.Ltime | log.Llongfile )
}