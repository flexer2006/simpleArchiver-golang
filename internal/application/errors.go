// Package application provides core utilities for error handling and panic recovery
// in CLI applications. It offers structured error logging, panic recovery with stack traces,
// and convenience wrappers for safe command execution.
package application

import (
	"errors"
	"log"
	"runtime"
)

// ErrEmptyPath is a sentinel error returned when required file path arguments are missing.
// Common usage includes validation of command-line parameters for file operations.
var ErrEmptyPath = errors.New("path to file is not specified")

// HandleError executes a error-returning function with centralized error logging.
// Parameters:
//   - fn: func() error - Function closure to execute that may return an error
//
// Returns:
//   - error: Original error if encountered, nil on successful execution
//
// Logs errors using logError before returning them to the caller.
func HandleError(fn func() error) error {
	if err := fn(); err != nil {
		logError(err)
		return err
	}
	return nil
}

// HandlePanic executes a function with panic recovery and stack trace logging.
// Parameters:
//   - fn: func() - Function closure to execute that might panic
//
// Returns:
//   - interface{}: Recovered panic value if panic occurred, nil otherwise
//
// Should be used as a top-level wrapper for critical code sections. Captures
// stack traces during panics using runtime.Stack for detailed diagnostics.
func HandlePanic(fn func()) (recovered interface{}) {
	defer func() {
		if r := recover(); r != nil {
			recovered = r
			logPanic(r)
		}
	}()
	fn()
	return nil
}

// logPanic logs panic details with stack traces. Uses runtime.Stack to capture
// up to 1024 bytes of stack information. Called automatically by HandlePanic.
func logPanic(recovered interface{}) {
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	log.Printf("PANIC: %v\nSTACK TRACE:\n%s", recovered, buf)
}

// logError provides standardized error logging format. Called automatically
// by HandleError for consistent error message formatting.
func logError(err error) {
	log.Printf("ERROR: %v", err)
}
