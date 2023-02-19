package null_logger

import (
	"fmt"
	"github.com/a-inacio/rosetta-logger-go/pkg/logger"
)

type NullLogger struct{}

func (l *NullLogger) Debug(args ...interface{})                 {}
func (l *NullLogger) Debugf(format string, args ...interface{}) {}
func (l *NullLogger) Info(args ...interface{})                  {}
func (l *NullLogger) Infof(format string, args ...interface{})  {}
func (l *NullLogger) Warn(args ...interface{})                  {}
func (l *NullLogger) Warnf(format string, args ...interface{})  {}
func (l *NullLogger) Error(args ...interface{})                 {}
func (l *NullLogger) Errorf(format string, args ...interface{}) {}
func (l *NullLogger) Fatal(args ...interface{}) {
	panic(fmt.Sprint(args...))
}
func (l *NullLogger) Fatalf(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}

func (l *NullLogger) WithFields(fields logger.Fields) logger.Logger {
	return l
}
