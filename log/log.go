package log

import (
	"context"
	"github.com/patrickmn/go-cache"
	"github/baishan-development-guizhou/golang-library/goid"
	"go.uber.org/zap"
	"time"
)

type loggerKey struct{}

var (
	global      ILogger
	loggerCache *cache.Cache
)

//Context returns context.Context and ILogger func.
//You should pass the Context,So that we can get the ILogger from the context
func Context(ctx context.Context) (context.Context, ILogger) {
	logger := ctx.Value(loggerKey{})
	if logger == nil {
		return Associate(ctx, Global())
	}
	return ctx, logger.(ILogger)
}

//Associate returns a copy of context.Context in which the ILogger associated
func Associate(ctx context.Context, logger ILogger) (context.Context, ILogger) {
	return context.WithValue(ctx, loggerKey{}, logger), logger
}

//Routine returns the ILogger bound to the current goroutine,
//we are not sure if the correct ILogger is returned,
//You'd better not use this func
func Routine() ILogger {
	if item, ok := loggerCache.Get(goid.GoID()); !ok {
		logger := Global()
		BindRoutine(logger)
		return logger
	} else {
		return item.(ILogger)
	}
}

//BindRoutine bind the ILogger to the current goroutine
func BindRoutine(logger ILogger) {
	loggerCache.Set(goid.GoID(), logger, cache.DefaultExpiration)
}

//Global returns global ILogger extends from zap.Logger
func Global() ILogger {
	return global
}

//ReplaceGlobal replace the global ILogger with the one passed by the parameter
func ReplaceGlobal(logger ILogger) {
	global = logger
}

func init() {
	global = &ZapLoggerAdapter{zap.NewNop().Sugar()}
	loggerCache = cache.New(30*time.Minute, 15*time.Minute)
}
