package logger

import (
	"os"
	"runtime"
	"sync"

	"github.com/rs/zerolog"
)

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

var once sync.Once
var log *zerolog.Logger

type loggerWrapper struct {
	log *zerolog.Logger
}

func GetLogger() Logger {

	// Run the function only once
	// Sigleton
	once.Do(func() {
		file, err := os.OpenFile("logger.json", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		// Setup Logger
		logger := zerolog.New(file).With().Timestamp().Logger()
		log = &logger
	})

	return &loggerWrapper{log}
}

func (l *loggerWrapper) Info(msg string) {
	file, line := getFileAndLine()
	l.log.Info().
		Str("file", file).
		Int("line", line).
		Msg(msg)
}

func (l *loggerWrapper) Warn(msg string) {
	file, line := getFileAndLine()
	l.log.Warn().
		Str("file", file).
		Int("line", line).
		Msg(msg)
}

func (l *loggerWrapper) Error(msg string) {
	file, line := getFileAndLine()
	l.log.Error().
		Str("file", file).
		Int("line", line).
		Msg(msg)
}

func (l *loggerWrapper) Debug(msg string) {
	file, line := getFileAndLine()
	l.log.Debug().
		Str("file", file).
		Int("line", line).
		Msg(msg)
}

func getFileAndLine() (string, int) {
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "unknown", 0
	}
	return file, line
}
