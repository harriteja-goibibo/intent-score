// package logger provides functionalities related to logging
//
// TODO: Validate Uber Zap (https://github.com/uber-go/zap)
// Zap supposed to provide better log throughput compared to normal logging.
// On manual benchmark ZAP had same performance as current logger. On top of that provides a better control on logging
package logger

import (
	"context"
	"log"
	"os"

	"github.com/pkg/errors"
)

// Logger struct that implements internal/Logger interface
type Logger struct {
	file    *os.File
	log     *log.Logger
	isDebug bool
}

// New returns new logger.
func New(name, filePath string) (*Logger, error) {
	return newLogger(name+" ", filePath)
}

// NewDebugLogger returns debug logger.
func NewDebugLogger(name, filePath string) (*Logger, error) {
	l, err := newLogger(name, filePath)
	l.isDebug = true
	return l, err
}

// newLogger creates a basic logger.
func newLogger(loggerName, filePath string) (*Logger, error) {
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, errors.Wrap(err, "newLogger")
	}
	l := Logger{
		file: f,
		log:  log.New(f, loggerName, log.Ldate|log.Ltime|log.Llongfile),
	}

	return &l, nil
}

// Close closes the file used for logging.
func (l *Logger) Close() {
	l.file.Close()
}

// Debug logs all the debug log given a logger.
func (l *Logger) Debug(ctx context.Context, bucket, handler, message, extraMessage string) {
	if l.isDebug {
		formattedLogMsg := makeLogMsg("DEBUG", bucket, handler, message, extraMessage)
		_ = l.log.Output(2, formattedLogMsg)
	}
}

// Info logs all the Info log given a logger.
func (l *Logger) Info(ctx context.Context, bucket, handler, message, extraMessage string) {
	formattedLogMsg := makeLogMsg("INFO", bucket, handler, message, extraMessage)
	_ = l.log.Output(2, formattedLogMsg)
}

// Error logs all the ERROR logs. It also pushes these errors to newrelic.
func (l *Logger) Error(ctx context.Context, bucket, handler, message, extraMessage string) {
	formattedLogMsg := makeLogMsg("ERROR", bucket, handler, message, extraMessage)
	_ = l.log.Output(2, formattedLogMsg)

}
