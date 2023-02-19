package logger

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	WithFields(fields Fields) Logger
}

type Fields map[string]interface{}

type Type int

const (
	NullLoggerType Type = iota
	StdLoggerType
	ZapAdapterLoggerType
	LogrusAdapterLoggerType
	GlogAdapterLoggerType
	ZerologAdapterLoggerType
)

type LevelType int

const (
	PanicLevel LevelType = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)
