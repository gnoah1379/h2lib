package idgen

import "time"

const (
	DefaultTimeUnit        = 1e7
	DefaultBitsTimestamp   = 40
	DefaultBitsSequenceIds = 10
	DefaultBitsServerIds   = 63 - DefaultBitsTimestamp - DefaultBitsSequenceIds
)

func DefaultStartTime() (time.Time, error) {
	return time.Parse(time.RFC3339, "2000-01-01T00:00:00Z")
}

func defaultGenerator(serverId int64) (g *generator) {
	g = &generator{
		sequence:      0,
		lastTime:      0,
		serverId:      serverId,
		bitsTimestamp: DefaultBitsTimestamp,
		bitsSequence:  DefaultBitsSequenceIds,
		bitsMachineId: DefaultBitsServerIds,
		timeUnit:      DefaultTimeUnit,
	}
	g.startTime, _ = DefaultStartTime()
	return g
}
