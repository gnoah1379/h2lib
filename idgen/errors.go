package idgen

import "errors"

var (
	InvalidConfig    = errors.New("sum of bitsTimestamp, bitsSequence, bitsMachineId must be equal 63")
	InvalidTimeUnit  = errors.New("timeunit exp must be great than 0 and less then 9")
	InvalidStartTime = errors.New("startTime must be before now")
	InvalidMachineID = errors.New("serverId must be great or equal than 0 and less then 8192")
)
