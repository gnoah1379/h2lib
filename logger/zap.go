package logger

import (
	"errors"
	"go.uber.org/zap"
)

func New(config zap.Config, fields ...zap.Field) (Logger, error) {
	log, err := config.Build()
	if err != nil {
		return nil, err
	}
	return log.With(fields...).Sugar(), err
}

func Wrap(l Logger, fields ...zap.Field) (Logger, error) {
	log, ok := l.(*zap.SugaredLogger)
	if !ok {
		return nil, errors.New("is not zap logger")
	}
	return log.Desugar().With(fields...).Sugar(), nil
}
