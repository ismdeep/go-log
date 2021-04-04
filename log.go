package log

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()
}

func Info(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Infow(msg, keysAndValues...)

}

func Debug(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Debugw(msg, keysAndValues...)
}

func Error(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Errorw(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Warnw(msg, keysAndValues...)
}

func Panic(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Panicw(msg, keysAndValues...)
}

func Fatal(msg string, keysAndValues ...interface{}) {
	logger.Sugar().Fatalw(msg, keysAndValues...)
}
