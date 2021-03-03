package echo_ext

import (
	glog "github.com/baishan-development-guizhou/golang-library/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"io"
)

type gLoggerAdapter struct {
	glog.Logger
}

func NewEchoLoggerWithGLogger(logger glog.Logger) echo.Logger {
	return &gLoggerAdapter{logger}
}

func (e *gLoggerAdapter) Output() io.Writer {
	glog.Warn("logger won't support it")
	return nil
}

func (e *gLoggerAdapter) SetOutput(io.Writer) {
	glog.Warn("logger won't support it")

}

func (e *gLoggerAdapter) Prefix() string {
	glog.Warn("logger won't support it")
	return ""
}

func (e *gLoggerAdapter) SetPrefix(string) {
	glog.Warn("logger won't support it")
}

func (e *gLoggerAdapter) Level() log.Lvl {
	glog.Warn("logger won't support it")
	return 0
}

func (e *gLoggerAdapter) SetLevel(log.Lvl) {
	glog.Warn("logger won't support it")
}

func (e *gLoggerAdapter) SetHeader(string) {
	glog.Warn("logger won't support it")
}

func (e *gLoggerAdapter) Print(i ...interface{}) { e.Logger.Info(i...) }

func (e *gLoggerAdapter) Printf(format string, args ...interface{}) {
	e.Logger.Infof(format, args...)
}

func (e *gLoggerAdapter) Printj(j log.JSON) { e.Logger.Info(j) }

func (e *gLoggerAdapter) Debugj(j log.JSON) { e.Logger.Debug(j) }

func (e *gLoggerAdapter) Infoj(j log.JSON) { e.Logger.Info(j) }

func (e *gLoggerAdapter) Warnj(j log.JSON) { e.Logger.Warn(j) }

func (e *gLoggerAdapter) Errorj(j log.JSON) { e.Logger.Error(j) }

func (e *gLoggerAdapter) Fatalj(j log.JSON) { e.Logger.Fatal(j) }

func (e *gLoggerAdapter) Panicj(j log.JSON) { e.Logger.Panic(j) }
