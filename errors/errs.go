package errors

import (
	"github.com/juju/errors"
)

// New is a drop in replacement for the standard libary errors module that records
// the location that the error is created.
func New(msg string) error {
	return errors.New(msg)
}

// Errorf creates a new annotated error and records the location that the
// error is created.  This should be a drop in replacement for fmt.Errorf.
func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args...)
}

// Trace adds the location of the Trace call to the stack.  The Cause of the
// resulting error is the same as the error parameter.  If the other error is
// nil, the result will be nil.
func Trace(err error) error {
	return errors.Trace(err)
}

// Annotate is used to add extra context to an existing error. The location of
// the Annotate call is recorded with the annotations. The file, line and
// function are also recorded.
func Annotate(err error, msg string) error {
	return errors.Annotate(err, msg)
}

// Annotatef is used to add extra context to an existing error. The location of
// the Annotate call is recorded with the annotations. The file, line and
// function are also recorded.
func Annotatef(err error, format string, args ...interface{}) error {
	return errors.Annotatef(err, format, args...)
}

// ErrorStack returns a string representation of the annotated error. If the
// error passed as the parameter is not an annotated error, the result is
// simply the result of the Error() method on that error.
//
// If the error is an annotated error, a multi-line string is returned where
// each line represents one entry in the annotation stack. The full filename
// from the call stack is used in the output.
func ErrorStack(err error) string {
	return errors.ErrorStack(err)
}

// Details returns information about the stack of errors wrapped by err, in
// the format:
//
// 	[{filename:99: error one} {otherfile:55: cause of error one}]
//
// This is a terse alternative to ErrorStack as it returns a single line.
func Details(err error) string {
	return errors.Details(err)
}

// Cause returns the cause of the given error.
// This will be either the original error.
// Cause is the usual way to diagnose errors that may have been wrapped by
// the other errors functions.
func Cause(err error) error {
	return errors.Cause(err)
}
