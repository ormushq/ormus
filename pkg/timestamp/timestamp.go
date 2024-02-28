package timestamp

import "time"

func Now() int64 {
	return time.Now().UnixMicro()
}

func Add(d time.Duration) int64 {
	return time.Now().Add(d).UnixMicro()
}
