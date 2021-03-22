package log

import (
	"go.uber.org/zap"
	"log"
)

//ZapLoggerAdapter adapter for Logger
type ZapLoggerAdapter struct {
	*zap.SugaredLogger
}

func (z *ZapLoggerAdapter) With(fields ...interface{}) Logger {
	return &ZapLoggerAdapter{z.SugaredLogger.With(fields...)}
}

func (z *ZapLoggerAdapter) Named(named string) Logger {
	return &ZapLoggerAdapter{z.SugaredLogger.Named(named)}
}

func (z *ZapLoggerAdapter) StdLogger() *log.Logger {
	return zap.NewStdLog(z.Desugar())
}
