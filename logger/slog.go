package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type slogLogger struct {
	appName string
	env     string
	lvl     Level
	slog    *slog.Logger
}

var _ Logger = (*slogLogger)(nil)

func NewSlogLogger(appName, env string, lvl Level) Logger {
	opts := SlogHandlerOptions{
		SlogOpts: slog.HandlerOptions{
			Level: levelToSlog(lvl),
		},
	}
	handler := NewSlogHandler(os.Stdout, opts)
	slogger := slog.New(handler).With(KeyAppName, appName, KeyEnv, env)

	return slogLogger{
		appName: appName,
		lvl:     lvl,
		slog:    slogger,
	}
}

func levelToSlog(lvl Level) slog.Level {
	return slog.Level(lvl)
}

func (s slogLogger) Debug(msg string, args ...any) {
	s.slog.Debug(msg, args...)
}

func (s slogLogger) Info(msg string, args ...any) {
	s.slog.Info(msg, args...)
}

func (s slogLogger) Warn(msg string, args ...any) {
	s.slog.Warn(msg, args...)
}

func (s slogLogger) Error(msg string, args ...any) {
	s.slog.Error(msg, args...)
}

func (s slogLogger) Fatal(msg string, args ...any) {
	s.slog.Error(msg, args...)
	os.Exit(1)
}

func (s slogLogger) Debugf(format string, args ...any) {
	s.Debug(fmt.Sprintf(format, args...))
}

func (s slogLogger) Infof(format string, args ...any) {
	s.Info(fmt.Sprintf(format, args...))
}

func (s slogLogger) Warnf(format string, args ...any) {
	s.Warn(fmt.Sprintf(format, args...))
}

func (s slogLogger) Errorf(format string, args ...any) {
	s.Error(fmt.Sprintf(format, args...))
}

func (s slogLogger) Fatalf(format string, args ...any) {
	s.Fatal(fmt.Sprintf(format, args...))
}

func (s slogLogger) With(args ...any) Logger {
	withLogger := s
	withLogger.slog = s.slog.With(args...)

	return withLogger
}

func (s slogLogger) WithComponent(component string) Logger {
	withComponentLogger := s
	withComponentLogger.slog = s.slog.With(ComponentName, component)

	return withComponentLogger
}
