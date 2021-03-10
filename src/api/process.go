package api

import (
	"github.com/gorilla/websocket"
)

type Pwss struct {
	Ws   *websocket.Conn
	Send chan interface{}
	Get  chan interface{}
}

// Prun ...
// idk
type Prun struct {
	Name string
	Cmd  string
}
