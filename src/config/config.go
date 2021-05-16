package config

import "time"

const (
	MaxSessionLiveTime = time.Duration(time.Minute * 60)
	MaxCatchRam        = int(30)
	MaxWSBufferSize    = int(1024 * 20)
	MaxAPIPayloadLen   = int(50)
	WSHandshakeTimeout = time.Duration(1000)
	ChannelBufferSize  = int(10)
	SessionIDLen       = int(64)
)
