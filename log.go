package log

import (
	"encoding/json"

	"github.com/ismdeep/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type configModel struct {
	Level            string   `json:"level"`
	Encoding         string   `json:"encoding"`
	OutputPaths      []string `json:"outputPaths"`
	ErrorOutputPaths []string `json:"errorOutputPaths"`
}

func loadConfig() configModel {
	var conf configModel
	if err := config.Load("log", &conf); err == nil {
		return conf
	}

	return configModel{
		Level:            "DEBUG",
		Encoding:         "console",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

var logger *zap.Logger

func init() {
	conf := loadConfig()
	raw, err := json.Marshal(conf)
	if err != nil {
		panic(err)
	}

	var cfg zap.Config
	if err := json.Unmarshal(raw, &cfg); err != nil {
		panic(err)
	}

	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	cfg.DisableCaller = true
	cfg.DisableStacktrace = false

	logger, err = cfg.Build()
	if err != nil {
		panic(err)
	}

	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
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

func String(key string, value string) zap.Field {
	return zap.String(key, value)
}

func FieldErr(err error) zap.Field {
	return zap.Error(err)
}

func Any(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}
