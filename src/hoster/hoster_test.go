package hoster

import (
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	ServerList["TOL"] = &List{Path: "./../../bds/testserver.bat"}
	Run()
	time.Sleep(time.Second * 2)
	ServerList["TOL"].EventChan <- "start"
	time.Sleep(time.Second * 4)
	ServerList["TOL"].EventChan <- "kill"
	time.Sleep(time.Second * 4)
	ServerList["TOL"].EventChan <- "start"
	time.Sleep(time.Second * 4)
	ServerList["TOL"].EventChan <- "restart"
	time.Sleep(time.Second * 4)
}
