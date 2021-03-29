package config

import "time"

const (
	MaxSessionLiveTime = int64(time.Minute * 60)
	MaxCatchRam        = int(30)
	MaxWSBufferSize    = int(1024 * 20)
)
