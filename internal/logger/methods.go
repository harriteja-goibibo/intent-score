package logger

import (
	"bytes"
)

const del = "\t"

// makeLogMsg formats log message.
func makeLogMsg(level, bucket, handler, message, extraMessage string) string {
	logBuffer := bytes.Buffer{}

	logBuffer.WriteString(level)
	logBuffer.WriteString(del)
	logBuffer.WriteString(handler)
	logBuffer.WriteString(del)
	logBuffer.WriteString(message)
	logBuffer.WriteString(del)
	logBuffer.WriteString(extraMessage)
	logBuffer.WriteString(del)
	logBuffer.WriteString(bucket)

	return logBuffer.String()
}
