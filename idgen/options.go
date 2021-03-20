package idgen

import (
	"math"
	"time"
)

type Option func(generator *generator) error

func WithStartTime(startTime time.Time) Option {
	return func(generator *generator) error {
		if startTime.After(time.Now()) {
			return InvalidStartTime
		}
		generator.startTime = startTime
		return nil
	}
}

func WithTimeUnit(exp int) Option {
	return func(generator *generator) error {
		if exp < 0 || exp > 9 {
			return InvalidTimeUnit
		}
		generator.timeUnit = int64(math.Pow(10, float64(exp)))
		return nil
	}
}

func WithBitsSequence(bits int) Option {
	return func(generator *generator) error {
		generator.bitsSequence = bits
		return nil
	}
}

func WithBitsMachineID(bits int) Option {
	return func(generator *generator) error {
		generator.bitsMachineId = bits
		return nil
	}
}

func WithBitsTimestamp(bits int) Option {
	return func(generator *generator) error {
		generator.bitsTimestamp = bits
		return nil
	}
}
