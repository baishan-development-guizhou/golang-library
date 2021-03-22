package log

import (
	"fmt"
	"github.com/baishan-development-guizhou/golang-library/ocommon/oid"
	"go.uber.org/zap/zapcore"
	"os"
)

type Logger interface {
	Debug(i ...interface{})
	Debugf(format string, args ...interface{})

	Info(i ...interface{})
	Infof(format string, args ...interface{})

	Warn(i ...interface{})
	Warnf(format string, args ...interface{})

	Error(i ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(i ...interface{})
	Fatalf(format string, args ...interface{})

	Panic(i ...interface{})
	Panicf(format string, args ...interface{})

	With(fields ...interface{}) Logger

	Named(named string) Logger

	Sync() error
}

type Level zapcore.Level

const (
	DebugLevel = Level(zapcore.DebugLevel)
	InfoLevel  = Level(zapcore.InfoLevel)
	WarnLevel  = Level(zapcore.WarnLevel)
	ErrorLevel = Level(zapcore.ErrorLevel)
	FatalLevel = Level(zapcore.FatalLevel)
	PanicLevel = Level(zapcore.PanicLevel)
)

type OutputEncoder func(cfg zapcore.EncoderConfig) zapcore.Encoder

var (
	JsonOutputEncoder    OutputEncoder = zapcore.NewJSONEncoder
	ConsoleOutputEncoder OutputEncoder = zapcore.NewConsoleEncoder
)

type CallerEncoder zapcore.CallerEncoder

var (
	FullCallerEncoder        CallerEncoder = zapcore.FullCallerEncoder
	FullRoutineCallerEncoder CallerEncoder = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(fmt.Sprintf("%d#%d %s", os.Getegid(), oid.ID(), caller.String()))
	}
	ShortCallerEncoder        CallerEncoder = zapcore.ShortCallerEncoder
	ShortRoutineCallerEncoder CallerEncoder = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(fmt.Sprintf("%d#%d %s", os.Getegid(), oid.ID(), caller.TrimmedPath()))
	}
)

type LevelEncoder zapcore.LevelEncoder

var (
	LowercaseLevelEncoder      LevelEncoder = zapcore.LowercaseLevelEncoder
	LowercaseColorLevelEncoder LevelEncoder = zapcore.LowercaseColorLevelEncoder
	CapitalLevelEncoder        LevelEncoder = zapcore.CapitalLevelEncoder
	CapitalColorLevelEncoder   LevelEncoder = zapcore.CapitalColorLevelEncoder
	BracketLevelEncoder        LevelEncoder = func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString("[" + level.String() + "]")
	}
)
