/*
Package logger defines API server's leveled operation and state change recording service.

The app is expected to log detailed internal processes and state changes. Each deployment has ability to specify
the level of log to be captured and provided to end user and/or automated log management facility.

Supported logging levels, ordered by increasing details of internal state change recording, are:
	- CRITICAL
    - ERROR
    - WARNING
    - NOTICE
    - INFO
    - DEBUG
Corresponding configuration package accepts one of these levels as a string literal.

For formatting specification see pkg/fmt package.
*/
package logger

// Logger specifies interface for leveled logger implementation.
type Logger interface {
	// Fatal logs fatal error without formatting.
	Fatal(args ...interface{})

	// Fatalf logs fatal error with formatting and placeholder constituents replacements.
	Fatalf(format string, args ...interface{})

	// Panic logs critical panic error without formatting.
	Panic(args ...interface{})

	// Panicf logs critical panic error with formatting and placeholder constituents replacements.
	Panicf(format string, args ...interface{})

	// Panic logs critical error without formatting.
	Critical(args ...interface{})

	// Panicf logs critical error  with formatting and placeholder constituents replacements.
	Criticalf(format string, args ...interface{})

	// Error logs regular error without formatting.
	Error(args ...interface{})

	// Errorf logs regular error with formatting and placeholder constituents replacements.
	Errorf(format string, args ...interface{})

	// Warning logs suspicious state situation without formatting.
	Warning(args ...interface{})

	// Warningf logs suspicious state situation with formatting and placeholder constituents replacements.
	Warningf(format string, args ...interface{})

	// Notice logs significant state change without formatting.
	Notice(args ...interface{})

	// Noticef logs significant state change with formatting and placeholder constituents replacements.
	Noticef(format string, args ...interface{})

	// Info logs common and regular state change without formatting.
	Info(args ...interface{})

	// Infof logs common and regular state change with formatting and placeholder constituents replacements.
	Infof(format string, args ...interface{})

	// Debug logs regular and detailed state change without formatting.
	Debug(args ...interface{})

	// Debugf logs regular and detailed state change with formatting and placeholder constituents replacements.
	Debugf(format string, args ...interface{})

	// Printf logs regular and detailed state change with formatting and placeholder constituents replacements.
	Printf(string, ...interface{})
}
