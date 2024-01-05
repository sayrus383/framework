package logger

import "strings"

type Level int

const ComponentName = "component"

const (
	LevelDebug Level = -4
	LevelInfo  Level = 0
	LevelWarn  Level = 4
	LevelError Level = 8
)

const TimeFormat = "2006-01-02 15:04:05.999"

const (
	KeyMessage = "message"
	KeyLevel   = "log_type"
	KeyTime    = "timestamp"
	KeyFile    = "file"
	KeyAppName = "module"
	KeyEnv     = "version"
)

//goland:noinspection SpellCheckingInspection
type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	Fatal(msg string, args ...any)

	/*
		Логирование по формату как fmt.Sprintf
	*/
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)

	/*
		Новый инстанс логгера, со смерженными значениями
	*/
	With(args ...any) Logger

	/*
		Новый инстанс логгера, для указания имени коммпонента
	*/
	WithComponent(component string) Logger
}

func StringToLevel(lvl string) Level {
	switch strings.ToUpper(lvl) {
	case "DEBUG":
		return LevelDebug
	case "INFO":
		return LevelInfo
	case "WARN":
		return LevelWarn
	case "ERROR":
		return LevelError
	default:
		return LevelDebug
	}
}
