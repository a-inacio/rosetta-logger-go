package rosetta

import (
	"github.com/a-inacio/rosetta-logger-go/internal/null_logger"
	"github.com/a-inacio/rosetta-logger-go/internal/std_logger"
	"github.com/a-inacio/rosetta-logger-go/internal/zap_logger"
	"github.com/a-inacio/rosetta-logger-go/pkg/logger"
)

func NewLogger(logType logger.Type) logger.Logger {
	if logType == logger.NullLoggerType || int(logType) >= len(_loggerFactories) {
		return &null_logger.NullLogger{}
	}
	return _loggerFactories[logType]()
}

var _loggerFactories = []func() logger.Logger{
	func() logger.Logger { return &null_logger.NullLogger{} },
	func() logger.Logger { return std_logger.NewStdLogger() },
	func() logger.Logger { return zap_logger.NewZapAdapterLogger() },
}
