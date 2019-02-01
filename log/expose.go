package log

import (
	"io"

	logger "github.com/Sirupsen/logrus"
	"github.com/juju/errors"
)

// SetOutput sets the standard logger output.
func SetOutput(out io.Writer) {
	logger.SetOutput(out)
}

// SetLevel sets the standard logger level.
func SetLevel(level logger.Level) {
	logger.SetLevel(level)
}

// DebugStack a error's detail at level Debug on the standard logger.
func DebugStack(err error) {
	logger.Debug(errors.ErrorStack(err))
}

// InfoStack a error's detail at level Info on the standard logger.
func InfoStack(err error) {
	logger.Info(errors.ErrorStack(err))
}

// WarnStack a error's detail at level Warn on the standard logger.
func WarnStack(err error) {
	logger.Warn(errors.ErrorStack(err))
}

// ErrorStack a error's detail at level Error on the standard logger.
func ErrorStack(err error) {
	logger.Error(errors.ErrorStack(err))
}

// FatalStack a error's detail at level Fatal on the standard logger.
func FatalStack(err error) {
	logger.Fatal(errors.ErrorStack(err))
}

// PanicStack logs a error's detail at level Panic on the standard logger.
func PanicStack(err error) {
	logger.Panic(errors.ErrorStack(err))
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	logger.Print(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	logger.Warning(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...interface{}) {
	logger.Warningf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}
