package utils

import (
	"bytes"
	"fmt"
	"log"
)

// Logger used to construct logger proccess
type Logger struct {
	IsDebug bool
}

// Info used to log all INFO's messages
func (l *Logger) Info(msg ...interface{}) {
	logger, buf := createNew()

	for _, m := range msg {
		msglog := fmt.Sprintf("INFO: %s", m)
		logger.Print(msglog)
		fmt.Print(buf)
	}
}

// Error used to catch any error messages
func (l *Logger) Error(msg ...interface{}) {
	logger, buf := createNew()

	for _, m := range msg {
		msglog := fmt.Sprintf("ERROR: %s", m)
		logger.Print(msglog)
		fmt.Print(buf)
	}
}

// Debug used to log all DEBUG's message.
// This method should be active only if current process
// started with -debug (true).
func (l *Logger) Debug(msg ...interface{}) {
	if l.IsDebug {
		logger, buf := createNew()

		for _, m := range msg {
			msglog := fmt.Sprintf("DEBUG: %s", m)
			logger.Print(msglog)
			fmt.Print(buf)
		}
	}
}

// LogBuilder --
func LogBuilder(isDebug bool) *Logger {
	builder := new(Logger)
	builder.IsDebug = isDebug
	return builder
}

func createNew() (*log.Logger, *bytes.Buffer) {
	var buf bytes.Buffer
	logger := log.New(&buf, "", log.Ldate|log.Ltime|log.LUTC)
	return logger, &buf
}
