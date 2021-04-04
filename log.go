package go_log

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()
}

func Info(msg string, fields ...interface{}) {
	logger.Sugar().Infow(msg, fields)

}

func Debug(msg string, fields ...interface{}) {
	logger.Sugar().Debugw(msg, fields)
}

func Error(msg string, fields ...interface{}) {
	logger.Sugar().Errorw(msg, fields)
}

func Warn(msg string, fields ...interface{}) {
	logger.Sugar().Warnw(msg, fields)
}

func Panic(msg string, fields ...interface{}) {
	logger.Sugar().Panicw(msg, fields)
}

func Fatal(msg string, fields ...interface{}) {
	logger.Sugar().Fatalw(msg, fields)
}
