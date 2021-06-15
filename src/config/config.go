package config

import (
	"regexp"
	"time"
)

const (
	MaxSessionLiveTime time.Duration = time.Minute * 60
	MaxCatchRam        int           = 30
	MaxWSBufferSize    int           = 1024 * 20
	MaxAPIPayloadLen   int           = 50
	WSHandshakeTimeout time.Duration = 1000
	ChannelSize        int           = 10
	SessionIDLen       int           = 64
	ServerIDLen        int           = 16
	ServerRootDir      string        = "./servers/"
	TestServerPath     string        = "C:/Users/kymcm/Documents/VSCode/gobds/bds/"
	TestServerCommand  string        = "bedrock_server.exe"
	// TestServerPath    string = "sh bds/"
	// TestServerCommand string = "testserver.sh"
)

var (
	ConsoleOutput = regexp.MustCompile(`(?P<level>[A-Z]*)\]\s(?P<output>.*)$`)
)
