package logs

import "log"

var LOG *log.Logger

func New(log *log.Logger) {
	LOG = log
}

func Error(message ...any) {
	LOG.Println("ERROR: ", message)
}

func Info(message ...any) {
	LOG.Println("INFO: ", message)
}

func Warning(message ...any) {
	LOG.Println("WARNING: ", message)
}
