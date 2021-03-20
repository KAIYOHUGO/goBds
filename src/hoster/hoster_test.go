package hoster

import (
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	ServerList["TOL"] = &List{Path: "C:\\Users\\kymcm\\Documents\\VSCode\\gobds\\bds\\testserver.bat"}
	Run()
	ServerList["TOL"].EventChan <- "start"
	time.Sleep(time.Second * 2)
	ServerList["TOL"].EventChan <- "kill"
	time.Sleep(time.Second * 1)
	ServerList["TOL"].EventChan <- "start"
	time.Sleep(time.Second * 1)
	ServerList["TOL"].EventChan <- "restart"
	time.Sleep(time.Second * 2)
}
