package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var (
	defaultTimeLayout = "02/Jan/2006:15:04:05 +0800"
)

func register(options options) ILogger {
	var (
		infoWriteSyncers  []zapcore.WriteSyncer
		errWriteSyncers   []zapcore.WriteSyncer
		cores             []zapcore.Core
		opts              []zap.Option
		infoHook, errHook io.Writer
	)
	if options.stdDisplay {
		infoWriteSyncers = append(infoWriteSyncers, zapcore.AddSync(os.Stdout))
		errWriteSyncers = append(errWriteSyncers, zapcore.AddSync(os.Stdout))
	}
	if options.outPath != "" || options.errPath != "" {
		if options.outPath != "" {
			infoHook = &lumberjack.Logger{
				Filename:   options.outPath,
				MaxSize:    options.maxSize,
				MaxBackups: options.maxBackups,
				MaxAge:     options.maxSize,
				Compress:   options.compress,
			}
			if options.outPath == options.errPath {
				errHook = infoHook
			}
		}
		if options.errPath != "" && options.outPath != options.errPath {
			errHook = &lumberjack.Logger{
				Filename:   options.errPath,
				MaxSize:    options.maxSize,
				MaxBackups: options.maxBackups,
				MaxAge:     options.maxSize,
				Compress:   options.compress,
			}
		}
		infoWriteSyncers = append(infoWriteSyncers, zapcore.AddSync(infoHook))
		errWriteSyncers = append(errWriteSyncers, zapcore.AddSync(errHook))
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	if options.timeLayout != "" {
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(options.timeLayout)
	} else {
		encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(defaultTimeLayout)
	}
	cores = []zapcore.Core{zapcore.NewCore(
		options.outputMode.parse()(encoderConfig),
		zapcore.NewMultiWriteSyncer(infoWriteSyncers...),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl < zapcore.WarnLevel
		}),
	), zapcore.NewCore(
		options.outputMode.parse()(encoderConfig),
		zapcore.NewMultiWriteSyncer(errWriteSyncers...),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.WarnLevel
		}),
	)}

	if options.caller {
		opts = append(opts, zap.AddCaller())
	}
	if options.stacktrace {
		opts = append(opts, zap.AddStacktrace(zapcore.WarnLevel))
	}

	return &ZapLoggerAdapter{zap.New(zapcore.NewTee(cores...), opts...).Sugar()}
}
