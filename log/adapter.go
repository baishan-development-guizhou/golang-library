package log

import (
	"go.uber.org/zap"
)

//ZapLoggerAdapter adapter for ILogger
type ZapLoggerAdapter struct {
	*zap.SugaredLogger
}

func (z *ZapLoggerAdapter) With(fields ...interface{}) ILogger {
	return &ZapLoggerAdapter{z.SugaredLogger.With(fields...)}
}
