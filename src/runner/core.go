package runner

import (
	"gobds/src/api"
	"gobds/src/msg"
)

var (
	// List ...
	// save server data
	List = make(map[string]*api.Plist)
	// Event ...
	// send start,stop,kill cmd
	Event = make(chan api.Prun)
	// Cmd ...
	// send cmd in bds
	Cmd = make(chan api.Prun)
)

// Run ...
// core func
func Run() {
	List["TOL"] = &api.Plist{
		Path: "C:\\Users\\kymcm\\Documents\\VSCode\\gobds\\bds\\bedrock_server.exe",
	}
	msg.Log("setup server")
	for _, e := range List {
		go e.Setup()
	}
	go listener()
}
