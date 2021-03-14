package wss

import (
	"gobds/src/usefull"
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
		client Client
		login  = false
		send   = make(chan interface{})
	)
	ws, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		usefull.Err("wss echo", err)
		return
	}
	defer func() {
		usefull.Wan("disconnect")
		ws.Close()
	}()
	usefull.Log("wss connect")
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
func echoPlugin(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		usefull.Err("wss echo", err)
		return
	}
	ws.Close()

}
