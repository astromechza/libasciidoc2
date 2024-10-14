// Package log provides some common logging utilities to gain insight into the operation of this library.
// It uses the standard built in slog (structured logging) package with levels however by default no logging is called
// unless enabled.

package log

import (
	"context"
	"fmt"
	"io"
	"log/slog"
)

// Default is the default logger called by this library. By default, this is all disabled, however you may replace this
// in order to see logs at whatever level you need them.
var Default *slog.Logger

// Reset will reset the Default logger back to it's disabled default
func Reset() {
	Default = slog.New(slog.NewTextHandler(io.Discard, &slog.HandlerOptions{Level: slog.LevelError}))
}

func init() {
	Reset()
}

func WithField(key, value string) *slog.Logger {
	return Default.With(slog.String(key, value))
}

// DebugEnabled may be used as a guard around expensive state dumping or encoding routines at debug level.
func DebugEnabled() bool {
	return Default.Enabled(context.Background(), slog.LevelDebug)
}

func InfoEnabled() bool {
	return Default.Enabled(context.Background(), slog.LevelInfo)
}

func Debug(message string) {
	Default.DebugContext(context.Background(), message)
}

// Debugf will format the log message with the args at debug level. It's a good idea to protect this with DebugEnabled
// if the args are expensive to calculate.
func Debugf(format string, args ...interface{}) {
	if DebugEnabled() {
		Debug(fmt.Sprintf(format, args...))
	}
}

func Info(message string) {
	Default.InfoContext(context.Background(), message)
}

// Infof will format the log message with the args at info level.
func Infof(format string, args ...interface{}) {
	if InfoEnabled() {
		Info(fmt.Sprintf(format, args...))
	}
}

func Warn(message string) {
	Default.WarnContext(context.Background(), message)
}

// Warnf will format the log message with the args at warning level.
func Warnf(format string, args ...interface{}) {
	if Default.Enabled(context.Background(), slog.LevelWarn) {
		Warn(fmt.Sprintf(format, args...))
	}
}

func Error(message string) {
	Default.ErrorContext(context.Background(), message)
}

// Errorf will format the log message with the args at error level.
func Errorf(format string, args ...interface{}) {
	if Default.Enabled(context.Background(), slog.LevelError) {
		Error(fmt.Sprintf(format, args...))
	}
}
