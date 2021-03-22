package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

type options struct {
	//file output config
	//------------------
	outputPath string
	//stderr output path
	errOutputPath string
	//Maximum file size
	maxSize int
	//Maximum number of file backups
	maxBackups int
	//Maximum file retention days
	maxAge int
	// Is compression enabled
	compress bool

	//-------------------
	//Is it displayed in standard output and standard error
	stdOutput bool
	//Output mode,the optional value is JsonOutputEncoder ConsoleOutputEncoder
	outPutEncoder OutputEncoder
	//Log level,the optional value is DebugLevel InfoLevel WarnLevel ErrorLevel FatalLevel PanicLevel
	level Level
	//Report callerEncoder
	callerEncoder CallerEncoder
	//Report levelEncoder
	levelEncoder LevelEncoder
	//Report Warn level stack trace
	stacktrace bool
	//time layout
	timeLayout string
	//init the named
	names []string
	//init with fields
	fields []interface{}
}

func Configure() *options {
	return &options{level: InfoLevel,
		maxSize:       30,
		maxBackups:    5,
		maxAge:        7,
		compress:      false,
		timeLayout:    "02/Jan/2006:15:04:05 +0800",
		levelEncoder:  BracketLevelEncoder,
		outPutEncoder: JsonOutputEncoder, callerEncoder: nil, stdOutput: true}
}

func (o *options) WithStacktrace(stacktrace bool) *options {
	o.stacktrace = stacktrace
	return o
}

func (o *options) WithErrOutputPath(errOutputPath string) *options {
	o.errOutputPath = errOutputPath
	return o
}

func (o *options) WithOutputPath(outputPath string) *options {
	o.outputPath = outputPath
	return o
}

func (o *options) WithMaxSize(maxSize int) *options {
	o.maxSize = maxSize
	return o
}

func (o *options) WithMaxBackups(maxBackups int) *options {
	o.maxBackups = maxBackups
	return o
}

func (o *options) WithTimeLayout(timeLayout string) *options {
	o.timeLayout = timeLayout
	return o
}

func (o *options) WithMaxAge(maxAge int) *options {
	o.maxAge = maxAge
	return o
}

func (o *options) WithStdOutput(stdOutput bool) *options {
	o.stdOutput = stdOutput
	return o
}

func (o *options) WithCompress(compress bool) *options {
	o.compress = compress
	return o
}

func (o *options) WithOutputEncoder(outputEncoder OutputEncoder) *options {
	o.outPutEncoder = outputEncoder
	return o
}

func (o *options) WithLevel(level Level) *options {
	o.level = level
	return o
}

func (o *options) WithCallerEncoder(callerEncoder CallerEncoder) *options {
	o.callerEncoder = callerEncoder
	return o
}

func (o *options) WithLevelEncoder(encoder LevelEncoder) *options {
	o.levelEncoder = encoder
	return o
}

func (o *options) WithNamed(names ...string) *options {
	if len(names) > 0 {
		o.names = names
	}
	return o
}

func (o *options) WithFields(fields ...interface{}) *options {
	switch len(fields) {
	case 0:
		break
	default:
		o.fields = fields
	}
	return o
}

func (o *options) Init() Logger {
	var (
		infoWriteSyncers  []zapcore.WriteSyncer
		errWriteSyncers   []zapcore.WriteSyncer
		cores             []zapcore.Core
		opts              []zap.Option
		infoHook, errHook io.Writer
		encoderConfig     = zap.NewProductionEncoderConfig()
	)
	if o.stdOutput {
		infoWriteSyncers = append(infoWriteSyncers, zapcore.AddSync(os.Stdout))
		errWriteSyncers = append(errWriteSyncers, zapcore.AddSync(os.Stderr))
	}
	if o.outputPath != "" || o.errOutputPath != "" {
		if o.outputPath != "" {
			infoHook = &lumberjack.Logger{
				Filename:   o.outputPath,
				MaxSize:    o.maxSize,
				MaxBackups: o.maxBackups,
				MaxAge:     o.maxSize,
				Compress:   o.compress,
			}
			if o.outputPath == o.errOutputPath {
				errHook = infoHook
			}
		}
		if o.errOutputPath != "" && o.outputPath != o.errOutputPath {
			errHook = &lumberjack.Logger{
				Filename:   o.errOutputPath,
				MaxSize:    o.maxSize,
				MaxBackups: o.maxBackups,
				MaxAge:     o.maxSize,
				Compress:   o.compress,
			}
		}
		infoWriteSyncers = append(infoWriteSyncers, zapcore.AddSync(infoHook))
		errWriteSyncers = append(errWriteSyncers, zapcore.AddSync(errHook))
	}

	if o.callerEncoder != nil {
		opts = append(opts, zap.AddCaller())
		encoderConfig.EncodeCaller = zapcore.CallerEncoder(o.callerEncoder)
	}

	encoderConfig.EncodeLevel = zapcore.LevelEncoder(o.levelEncoder)
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(o.timeLayout)
	//fix #15
	encoderConfig.ConsoleSeparator = " "
	cores = []zapcore.Core{zapcore.NewCore(
		o.outPutEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(infoWriteSyncers...),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.Level(o.level) && lvl < zapcore.WarnLevel
		}),
	), zapcore.NewCore(
		o.outPutEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(errWriteSyncers...),
		zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.Level(o.level) && lvl >= zapcore.WarnLevel
		}),
	)}

	if o.stacktrace {
		opts = append(opts, zap.AddStacktrace(zapcore.WarnLevel))
	}
	zapSugarLogger := zap.New(zapcore.NewTee(cores...), opts...).Sugar()
	if o.names != nil {
		for _, named := range o.names {
			zapSugarLogger = zapSugarLogger.Named(named)
		}
	}

	if o.fields != nil {
		zapSugarLogger = zapSugarLogger.With(o.fields...)
	}

	return &ZapLoggerAdapter{zapSugarLogger}
}
