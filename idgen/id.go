package idgen

import "time"

type ID struct {
	Id         int64
	Timestamp  int64
	SequenceId int64
	ServerId   int64
	Time       time.Time
}
