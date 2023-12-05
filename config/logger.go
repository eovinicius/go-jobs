package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func (l *Logger) Flags() int {
	return log.Ldate | log.Ltime | log.Lshortfile
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

func (logger *Logger) Debug(value ...interface{}) {
	logger.debug.Println(value...)
}
func (logger *Logger) Info(value ...interface{}) {
	logger.info.Println(value...)
}
func (logger *Logger) Warning(value ...interface{}) {
	logger.warning.Println(value...)
}
func (logger *Logger) Err(value ...interface{}) {
	logger.err.Println(value...)
}

func (logger *Logger) DebugF(format string, value ...interface{}) {
	logger.debug.Printf(format, value...)
}
func (logger *Logger) InfoF(format string, value ...interface{}) {
	logger.info.Printf(format, value...)
}
func (logger *Logger) WarningF(format string, value ...interface{}) {
	logger.warning.Printf(format, value...)
}
func (logger *Logger) Errf(format string, value ...interface{}) {
	logger.err.Printf(format, value...)
}
