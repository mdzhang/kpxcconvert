package logger

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("kpxcconvert")
var prefix = fmt.Sprintf("%s ▶ ", color.GreenString("kpxcconvert"))

var format = logging.MustStringFormatter(
	prefix + `%{level:.4s} %{shortfunc} %{id:03x} %{message}`,
)

// Debug wraps regular logger with useful prefix, colors, and formatting
func Debug(f string, a ...interface{}) {
	log.Info(fmt.Sprintf(f, a...))
}

// Info wraps regular logger with useful prefix, colors, and formatting
func Info(f string, a ...interface{}) {
	log.Info(fmt.Sprintf(f, a...))
}

// Warn wraps regular logger with useful prefix, colors, and formatting
func Warn(f string, a ...interface{}) {
	log.Warning(fmt.Sprintf(f, a...))
}

// Error wraps regular logger with useful prefix, colors, and formatting
func Error(f string, a ...interface{}) {
	log.Error(fmt.Sprintf(f, a...))
}

// Init sets the log formatting
func Init() {
	logging.SetFormatter(format)
}
