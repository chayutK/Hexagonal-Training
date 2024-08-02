package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = ""

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(msg string, field ...zapcore.Field) {
	log.Info(msg, field...)
}

func Debug(msg string, field ...zapcore.Field) {
	log.Debug(msg, field...)
}

func Error(message interface{}, field ...zapcore.Field) {
	switch v := message.(type) {
	case error:
		log.Error(v.Error(), field...)
	case string:
		log.Error(v, field...)
	}
	// log.Error(msg, field...)
}
