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
	s := make(chan interface{})
	ServerList["TOL"].Broadcast.Add(s)
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
