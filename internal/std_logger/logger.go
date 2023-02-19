package std_logger

import (
	"fmt"
	"github.com/a-inacio/rosetta-logger-go/pkg/logger"
	"log"
	"os"
)

type StdLogger struct {
	logger *log.Logger
}

func NewStdLogger() *StdLogger {
	return &StdLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (l *StdLogger) Debug(v ...interface{}) {
	l.logger.Println("[DEBUG]", fmt.Sprint(v...))
}

func (l *StdLogger) Debugf(format string, v ...interface{}) {
	l.logger.Printf("[DEBUG] "+format, v...)
}

func (l *StdLogger) Info(v ...interface{}) {
	l.logger.Println("[INFO]", fmt.Sprint(v...))
}

func (l *StdLogger) Infof(format string, v ...interface{}) {
	l.logger.Printf("[INFO] "+format, v...)
}

func (l *StdLogger) Warn(v ...interface{}) {
	l.logger.Println("[WARN]", fmt.Sprint(v...))
}

func (l *StdLogger) Warnf(format string, v ...interface{}) {
	l.logger.Printf("[WARN] "+format, v...)
}

func (l *StdLogger) Error(v ...interface{}) {
	l.logger.Println("[ERROR]", fmt.Sprint(v...))
}

func (l *StdLogger) Errorf(format string, v ...interface{}) {
	l.logger.Printf("[ERROR] "+format, v...)
}

func (l *StdLogger) Fatal(v ...interface{}) {
	l.logger.Panicf("[FATAL]:", fmt.Sprint(v...))
}

func (l *StdLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Panicf("[FATAL]: "+format, args...)
}

func (l *StdLogger) WithFields(fields logger.Fields) logger.Logger {
	logger := log.New(l.logger.Writer(), l.logger.Prefix(), l.logger.Flags())
	logger.SetPrefix(l.logger.Prefix() + formatFields(fields))
	return &StdLogger{logger}
}

func formatFields(fields logger.Fields) string {
	formatted := ""
	for key, value := range fields {
		formatted += " " + key + "=" + fmt.Sprint(value)
	}
	return formatted
}
