package zap_logger

import (
	"github.com/a-inacio/rosetta-logger-go/pkg/logger"
	"go.uber.org/zap"
)

type ZapAdapterLogger struct {
	Logger *zap.SugaredLogger
}

func NewZapAdapterLogger() *ZapAdapterLogger {
	return &ZapAdapterLogger{
		Logger: zap.S(),
	}
}

func (l *ZapAdapterLogger) Debugf(msg string, args ...interface{}) {
	l.Logger.Debugf(msg, args...)
}

func (l *ZapAdapterLogger) Debug(args ...interface{}) {
	l.Logger.Debug(args...)
}

func (l *ZapAdapterLogger) Infof(msg string, args ...interface{}) {
	l.Logger.Infof(msg, args...)
}

func (l *ZapAdapterLogger) Info(args ...interface{}) {
	l.Logger.Info(args...)
}

func (l *ZapAdapterLogger) Warnf(msg string, args ...interface{}) {
	l.Logger.Warnf(msg, args...)
}

func (l *ZapAdapterLogger) Warn(args ...interface{}) {
	l.Logger.Warn(args...)
}

func (l *ZapAdapterLogger) Errorf(msg string, args ...interface{}) {
	l.Logger.Errorf(msg, args...)
}

func (l *ZapAdapterLogger) Error(args ...interface{}) {
	l.Logger.Error(args...)
}

func (l *ZapAdapterLogger) Fatalf(msg string, args ...interface{}) {
	l.Logger.Fatalf(msg, args...)
}

func (l *ZapAdapterLogger) Fatal(args ...interface{}) {
	l.Logger.Fatal(args...)
}

func (l *ZapAdapterLogger) WithFields(fields logger.Fields) logger.Logger {
	// Convert fields to zap.Fields
	zapFields := make([]interface{}, 0, len(fields)*2)
	for k, v := range fields {
		zapFields = append(zapFields, k, v)
	}
	// Create a new SugaredLogger instance with the added fields
	return &ZapAdapterLogger{l.Logger.With(zapFields...)}
}
