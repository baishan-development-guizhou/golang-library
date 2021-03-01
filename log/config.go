package log

type options struct {
	//file output config
	//------------------
	outPath string
	//stderr output path
	errPath string
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
	stdDisplay bool
	//Output mode,the optional value is JsonMode ConsoleMode
	outputMode OutPutMode
	//Log level,the optional value is DebugLevel InfoLevel WarnLevel ErrorLevel
	level Level
	//Report file line number
	caller bool
	//Report Warn level stack trace
	stacktrace bool
	//time layout
	timeLayout string
}

func (o *options) WithOutPath(outPath string) *options {
	o.outPath = outPath
	return o
}
func (o *options) WithStacktrace(stacktrace bool) *options {
	o.stacktrace = stacktrace
	return o
}
func (o *options) WithErrPath(errPath string) *options {
	o.errPath = errPath
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
func (o *options) WithStdDisplay(stdDisplay bool) *options {
	o.stdDisplay = stdDisplay
	return o
}

func (o *options) WithCompress(compress bool) *options {
	o.compress = compress
	return o
}

func (o *options) WithOutputMode(mode OutPutMode) *options {
	o.outputMode = mode
	return o
}

func (o *options) WithLevel(level Level) *options {
	o.level = level
	return o
}

func (o *options) WithCaller(caller bool) *options {
	o.caller = caller
	return o
}

func (o *options) Init() ILogger {
	return register(*o)
}

func Configure() *options {
	return &options{level: InfoLevel, maxSize: 30, maxBackups: 5, maxAge: 7, outputMode: ConsoleMode, caller: true, compress: false, stdDisplay: true}
}
