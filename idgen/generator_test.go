package idgen

import (
	"encoding/json"
	"runtime"
	"testing"
	"time"
)

var gen Generator
var startTime int64
var serverId int64

func TestOnce(t *testing.T) {
	var sleepTime int64 = 50
	time.Sleep(time.Duration(sleepTime) * 10 * time.Millisecond)
	id := gen.NewID()

	if id < 0 {
		t.Errorf("id nagative: %d", id)
	}
	info := gen.Decode(id)
	if info.Timestamp < sleepTime || info.Timestamp > sleepTime+1 {
		t.Errorf("unexpected actual time: %d != sleep time: %d", info.Timestamp, sleepTime)
	}

	if info.SequenceId != 0 {
		t.Errorf("unexpected sequence: %d", info.SequenceId)
	}

	if info.ServerId != serverId {
		t.Errorf("unexpected infrastructure id: %d", info.ServerId)
	}

	t.Log("generate id:", id)
	infoJson, _ := json.Marshal(info)
	t.Log("decode:", string(infoJson))
}

func Test10Sec(t *testing.T) {
	var numID uint32
	var lastID int64
	var maxSeqId int64
	initial := currentTime()
	current := initial
	for current-initial < 1000 {
		id := gen.NewID()
		if id < 0 {
			t.Errorf("id nagative: %d", id)
		}
		info := gen.Decode(id)
		numID++
		if id <= lastID {
			t.Fatal("duplicated id")
		}
		lastID = id

		current = currentTime()

		overtime := startTime + info.Timestamp - current
		if overtime > 0 {
			t.Log(numID)
			t.Log(overtime)
			t.Fatalf("unexpected overtime: %d", overtime)
		}

		if maxSeqId < info.SequenceId {
			maxSeqId = info.SequenceId
		}

		if info.ServerId != serverId {
			t.Errorf("unexpected infrastructure id: %d", info.ServerId)
		}
	}

	if maxSeqId != 1<<DefaultBitsSequenceIds-1 {
		t.Errorf("unexpected max sequence: %d", maxSeqId)
	}
	t.Log("number of id:", numID)
}

func TestParallel(t *testing.T) {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	t.Log("number of cpu:", numCPU)

	consumer := make(chan int64)

	const numID = 10000
	generate := func() {
		for i := 0; i < numID; i++ {
			consumer <- gen.NewID()
		}
	}

	const numGenerator = 10
	for i := 0; i < numGenerator; i++ {
		go generate()
	}

	set := make(map[int64]struct{})
	for i := 0; i < numID*numGenerator; i++ {
		id := <-consumer
		if _, ok := set[id]; ok {
			t.Fatal("duplicated id")
		}
		set[id] = struct{}{}
	}
	t.Log("number of id:", len(set))
}

func init() {
	var err error
	serverId = 0
	t := time.Now()
	startTime = t.UTC().UnixNano() / DefaultTimeUnit
	gen, err = New(int(serverId), WithStartTime(t))
	if err != nil {
		panic("can't not created gen")
	}
}

func currentTime() int64 {
	return time.Now().UTC().UnixNano() / DefaultTimeUnit
}
