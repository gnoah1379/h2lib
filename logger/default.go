package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Development(fields ...zap.Field) (Logger, error) {
	return New(DevelopmentConfig(), fields...)
}

func Production(fields ...zap.Field) (Logger, error) {
	return New(ProductionConfig(), fields...)
}

func ProductionConfig() zap.Config {
	conf := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		Encoding:         "json",
		EncoderConfig: zapcore.EncoderConfig{
			// Keys can be anything except the empty string.
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stack",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
	return conf
}

func DevelopmentConfig() zap.Config {
	conf := zap.Config{
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		Development:      true,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		Encoding:         "console",
		EncoderConfig: zapcore.EncoderConfig{
			// Keys can be anything except the empty string.
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stack",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}
	return conf
}
