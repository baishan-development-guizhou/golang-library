package log

import (
	"context"
	"go.uber.org/zap"
	"sync"
)

type loggerKey struct{}

var (
	_global   Logger
	_globalMu sync.RWMutex
)

//C returns Logger from context.Context func.
//You should pass the Context,So that we can get the Logger from the context
func C(ctx context.Context) Logger {
	var item interface{}
	if ctx != nil {
		item = ctx.Value(loggerKey{})
		if item == nil {
			item = G()
		}
	} else {
		item = G()
	}
	return item.(Logger)
}

//AssociateC returns a copy of context.Context in which the Logger associated
func AssociateC(ctx context.Context, logger Logger) (context.Context, Logger) {
	return context.WithValue(ctx, loggerKey{}, logger), logger
}

//G returns _global Logger extends from zap.Logger
func G() Logger {
	_globalMu.RLock()
	defer _globalMu.RUnlock()
	return _global
}

//ReplaceG replace the _global logger with the one passed by the parameter
func ReplaceG(logger Logger) {
	_globalMu.Lock()
	_global = logger
	_globalMu.Unlock()
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the _global logger.
func Debug(args ...interface{}) {
	_globalMu.RLock()
	_global.Debug(args...)
	_globalMu.RUnlock()
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(format string, args ...interface{}) {
	_globalMu.RLock()
	_global.Debugf(format, args...)
	_globalMu.RUnlock()
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the _global logger.
func Info(args ...interface{}) {
	_globalMu.RLock()
	_global.Info(args...)
	defer _globalMu.RUnlock()
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(format string, args ...interface{}) {
	_globalMu.RLock()
	_global.Infof(format, args...)
	_globalMu.RUnlock()
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the _global logger.
func Warn(args ...interface{}) {
	_globalMu.RLock()
	_global.Warn(args...)
	_globalMu.RUnlock()
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(format string, args ...interface{}) {
	_globalMu.RLock()
	_global.Warnf(format, args...)
	_globalMu.RUnlock()
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the _global logger.
func Error(args ...interface{}) {
	_globalMu.RLock()
	_global.Error(args...)
	_globalMu.RUnlock()
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(format string, args ...interface{}) {
	_globalMu.RLock()
	_global.Errorf(format, args...)
	_globalMu.RUnlock()
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the _global logger.
func Fatal(args ...interface{}) {
	_globalMu.RLock()
	_global.Fatal(args...)
	_globalMu.RUnlock()
}

// Fatalf uses fmt.Sprintf to log a templated message.
func Fatalf(format string, args ...interface{}) {
	_globalMu.RLock()
	_global.Fatalf(format, args...)
	_globalMu.RUnlock()
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the _global logger.
func Panic(args ...interface{}) {
	_globalMu.RLock()
	_global.Panic(args...)
	_globalMu.RUnlock()
}

// Panicf uses fmt.Sprintf to log a templated message.
func Panicf(format string, args ...interface{}) {
	_globalMu.RLock()
	_global.Panicf(format, args...)
	_globalMu.RUnlock()
}

// With adds a variadic number of fields to the logging context.
func With(fields ...interface{}) Logger {
	_globalMu.RLock()
	defer _globalMu.RUnlock()
	return _global.With(fields...)
}

// Named adds a sub-scope to the logger's name.
func Named(named string) Logger {
	_globalMu.RLock()
	defer _globalMu.RUnlock()
	return _global.Named(named)
}

// Sync flushing any buffered log entries.
//Applications should take care to call Sync before exiting.
func Sync() error {
	_globalMu.RLock()
	defer _globalMu.RUnlock()
	return _global.Sync()
}

func init() {
	_global = &ZapLoggerAdapter{zap.NewNop().Sugar()}
}
