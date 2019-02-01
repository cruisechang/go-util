package log

import (
	"io"

	"github.com/Sirupsen/logrus"
	"github.com/juju/errors"
)

// These are the different logging levels. You can set the logging level to log
// on your instance of logger.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel = logrus.PanicLevel
	// FatalLevel level. Logs and then calls `os.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel = logrus.FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel = logrus.ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel = logrus.WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel = logrus.InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel = logrus.DebugLevel
)

// Logger is a wrapper of logrus logger type.
type Logger struct {
	logger *logrus.Logger
}

// New returns a Logger instance with logrus default settings.
func New() *Logger {
	return &Logger{
		logger: logrus.New(),
	}
}

// SetOutput sets the logger output.
func (l *Logger) SetOutput(out io.Writer) {
	l.logger.Out = out
}

// SetLevel sets the logger level.
func (l *Logger) SetLevel(level logrus.Level) {
	l.logger.Level = level
}

// DebugStack log the debug level error detail.
func (l *Logger) DebugStack(err error) {
	l.logger.Debug(errors.ErrorStack(err))
}

// InfoStack log the info level error detail.
func (l *Logger) InfoStack(err error) {
	l.logger.Info(errors.ErrorStack(err))
}

// WarnStack log the warn level error detail.
func (l *Logger) WarnStack(err error) {
	l.logger.Warn(errors.ErrorStack(err))
}

// ErrorStack log the error level error detail.
func (l *Logger) ErrorStack(err error) {
	l.logger.Error(errors.ErrorStack(err))
}

// FatalStack log the fatal level error detail.
func (l *Logger) FatalStack(err error) {
	l.logger.Fatal(errors.ErrorStack(err))
}

// PanicStack log the panic level error detail.
func (l *Logger) PanicStack(err error) {
	l.logger.Panic(errors.ErrorStack(err))
}

// Debugf log the debug level message.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

// Infof log the info level message.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

// Warnf log the warn level message.
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

// Errorf log the error level message.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

// Fatalf log the fatal level message then fatal.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

// Panicf log the panic level message then panic.
func (l *Logger) Panicf(format string, args ...interface{}) {
	l.logger.Panicf(format, args...)
}

// Debug log the debug level message.
func (l *Logger) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

// Info log the info level message.
func (l *Logger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

// Warn log the warn level message.
func (l *Logger) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

// Error log the error level message.
func (l *Logger) Error(args ...interface{}) {
	l.logger.Error(args...)
}

// Fatal log the fatal level message then fatal.
func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal(args...)
}

// Panic log the panic level message then panic.
func (l *Logger) Panic(args ...interface{}) {
	l.logger.Panic(args...)
}
