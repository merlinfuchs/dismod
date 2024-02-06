package disutil

import (
	"fmt"
	"log"
)

type LogLevel int

const (
	LogDebug LogLevel = iota
	LogInfo
	LogWarn
	LogError
	LogFatal
)

func (l LogLevel) String() string {
	switch l {
	case LogDebug:
		return "DEBUG"
	case LogInfo:
		return "INFO"
	case LogWarn:
		return "WARN"
	case LogError:
		return "ERROR"
	case LogFatal:
		return "FATAL"
	}
	return "UNKNOWN"
}

type Logger func(level LogLevel, format string, a ...interface{})

func DefaultLogger(level LogLevel, format string, a ...interface{}) {
	var prefix string
	switch level {
	case LogDebug:
		prefix = "DEBUG"
	case LogInfo:
		prefix = "INFO"
	case LogWarn:
		prefix = "WARN"
	case LogError:
		prefix = "ERROR"
	case LogFatal:
		prefix = "FATAL"
	}
	log.Printf("[%s] %s", prefix, fmt.Sprintf(format, a...))
}
