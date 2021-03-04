package wss

import (
	"gobds/src/api"
	"gobds/src/msg"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{
	HandshakeTimeout: 0,
	ReadBufferSize:   0,
	WriteBufferSize:  0,
	WriteBufferPool:  nil,
	Subprotocols:     []string{},
	Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
	},
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: false,
}

func echoAPI(w http.ResponseWriter, r *http.Request) {
	var (
		client api.Gclient
		login  = false
		send   = make(chan interface{})
	)
	ws, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		msg.Err("wss echo", err)
		return
	}
	defer func() {
		msg.Wan("disconnect")
		ws.Close()
	}()
	msg.Log("wss connect")
	go func() {
		for el := range send {
			ws.WriteJSON(el)
		}
	}()
	for {
		if err = ws.ReadJSON(&client); err != nil {
			ws.WriteJSON(fail)
			break
		}
		send <- get(client, &login)
	}
	return
}
func echoCmd(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		msg.Err("wss echo", err)
		return
	}
}
func echoPlugin(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		msg.Err("wss echo", err)
		return
	}
}
