package idgen

import (
	"sync"
	"time"
)

type Generator interface {
	NewID() int64
	Decode(id int64) ID
	MaxSequenceId() int64
	MaxServerId() int64
	TimestampLimit() int64
	SequenceIdLimit() int64
	ServerIdLimit() int64
	TimeUnit() int64
}

type generator struct {
	startTime     time.Time
	serverId      int64
	sequence      int64
	lastTime      int64
	bitsTimestamp int
	bitsSequence  int
	bitsMachineId int
	timeUnit      int64
	mu            sync.RWMutex
}

func New(serverId int, opts ...Option) (Generator, error) {
	server := int64(serverId)
	g := defaultGenerator(server)
	for _, opt := range opts {
		err := opt(g)
		if err != nil {
			return nil, err
		}
	}
	if g.bitsTimestamp+g.bitsSequence+g.bitsMachineId != 63 {
		return nil, InvalidConfig
	}
	if serverId < 0 || server >= g.ServerIdLimit() {
		return nil, InvalidMachineID
	}
	return g, nil
}

func (g *generator) NewID() int64 {
	g.mu.Lock()
	now := g.timeNow()
	if now >= g.TimestampLimit() {
		return -1
	}
	seq := g.nextSequenceId(now)
	defer func() {
		if seq == g.SequenceIdLimit()-1 {
			sleepTime := time.Duration(now + g.TimeUnit() - int64(time.Since(g.startTime))%g.TimeUnit())
			time.Sleep(sleepTime)
		}
		g.lastTime = now
		g.mu.Unlock()
	}()
	return g.newID(now, seq)
}

func (g *generator) Decode(id int64) ID {
	info := ID{
		Id:         id,
		Timestamp:  id >> (g.bitsSequence + g.bitsMachineId),
		SequenceId: id & (g.MaxSequenceId() << int64(g.bitsMachineId)) >> g.bitsMachineId,
		ServerId:   id & g.MaxServerId(),
	}
	info.Time = g.startTime.Add(time.Duration(info.Timestamp * g.TimeUnit()))
	return info
}

func (g *generator) MaxSequenceId() int64 {
	return 1<<g.bitsSequence - 1
}

func (g *generator) MaxServerId() int64 {
	return 1<<g.bitsMachineId - 1
}

func (g *generator) TimestampLimit() int64 {
	return 1 << g.bitsTimestamp
}

func (g *generator) SequenceIdLimit() int64 {
	return 1 << g.bitsSequence
}

func (g *generator) ServerIdLimit() int64 {
	return 1 << g.bitsMachineId
}

func (g *generator) TimeUnit() int64 {
	return g.timeUnit
}

func (g *generator) newID(now int64, seq int64) int64 {
	return now<<(g.bitsMachineId+g.bitsSequence) |
		seq<<g.bitsMachineId |
		g.serverId
}

func (g *generator) nextSequenceId(now int64) int64 {
	if now == g.lastTime {
		g.sequence++
		return g.sequence % g.SequenceIdLimit()
	} else {
		g.sequence = 0
		return 0
	}
}

func (g *generator) timeNow() (now int64) {
	return time.Since(g.startTime).Nanoseconds() / g.TimeUnit()
}
