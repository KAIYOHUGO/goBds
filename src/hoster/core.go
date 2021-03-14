package hoster

import (
	hoster "gobds/src/api"
	"gobds/src/usefull"
)

var (
	// ServerList ...
	// save server data
	ServerList = make(map[string]*List)
	// Event ...
	// send start,stop,kill cmd
	Event = make(chan hoster.Prun)
	// Cmd ...
	// send cmd in bds
	Cmd = make(chan hoster.Prun)
)

// Run ...
// core func
func Run() {
	ServerList["TOL"] = &List{
		Path: "C:\\Users\\kymcm\\Documents\\VSCode\\gobds\\bds\\bedrock_server.exe",
	}
	usefull.Log("setup server")
	for _, e := range ServerList {
		go e.Setup()
	}
	go listener()
}
