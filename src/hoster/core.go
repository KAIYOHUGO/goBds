package hoster

import (
	"gobds/src/usefull"
)

var (
	// ServerList ...
	// save server data
	ServerList = make(map[string]*List)
	// Event ...
	// send start,stop,kill cmd
	// Event = make(chan string)
	// Cmd ...
	// send cmd in bds
	// Cmd = make(chan string)
)

// Run ...
// core func
func Run() {
	usefull.Log("setup server")
	for _, e := range ServerList {
		e.init()
	}
}
