package internal

import "context"

// Logger is the main interface used for logging.
type Logger interface {

	// Debug logs all the DEBUG level log messages.
	// The implementer has to make sure that this is enabled only for development
	// or if debug is set to true; else the message should be ignored.
	Debug(ctx context.Context, bucket, handler, message, extraMessage string)

	// Info logs all INFO level log messages.
	Info(ctx context.Context, bucket, handler, message, extraMessage string)

	// Error logs all ERROR level log messages.
	Error(ctx context.Context, bucket, handler, message, extraMessage string)

	// Close is used for cleaning up opened resources
	// like files, connections or buffers.
	Close()
}
