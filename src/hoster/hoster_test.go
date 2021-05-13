package hoster

import (
	"gobds/src/utils"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	// ServerList["TOL"] = &List{Path: "./../../bds/testserver.bat"}
	ServerList["TOL"] = &List{Path: "./../../bds/bedrock_server.exe"}
	Run()
	s, l := ServerList["TOL"].Broadcast.New()
	defer ServerList["TOL"].Broadcast.Close(l)
	go func() {
		for i := range s {
			utils.Log("broadcast:" + i.(string))
		}
	}()
	// time.Sleep(time.Second * 2)
	ServerList["TOL"].CmdChan <- "$start"
	time.Sleep(time.Second * 2)
	ServerList["TOL"].CmdChan <- "$kill"
	time.Sleep(time.Second * 10)
	ServerList["TOL"].CmdChan <- "$start"
	// time.Sleep(time.Second * 10)
	time.Sleep(time.Second * 20)
}
