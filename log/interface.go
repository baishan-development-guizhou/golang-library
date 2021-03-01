package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ILogger interface {
	Debug(i ...interface{})
	Debugf(format string, args ...interface{})

	Info(i ...interface{})
	Infof(format string, args ...interface{})

	Warn(i ...interface{})
	Warnf(format string, args ...interface{})

	Error(i ...interface{})
	Errorf(format string, args ...interface{})

	With(fields ...interface{}) ILogger
}

type Level zapcore.Level

func (l Level) parse() zap.AtomicLevel {
	return zap.NewAtomicLevelAt(zapcore.Level(l))
}

const (
	DebugLevel = Level(zapcore.DebugLevel)
	InfoLevel  = Level(zapcore.InfoLevel)
	WarnLevel  = Level(zapcore.WarnLevel)
	ErrorLevel = Level(zapcore.ErrorLevel)
)

type OutPutMode uint

func (o OutPutMode) parse() func(cfg zapcore.EncoderConfig) zapcore.Encoder {
	switch o {
	case JsonMode:
		return zapcore.NewJSONEncoder
	case ConsoleMode:
		return zapcore.NewConsoleEncoder
	default:
		return zapcore.NewConsoleEncoder
	}
}

const (
	JsonMode OutPutMode = iota
	ConsoleMode
)
