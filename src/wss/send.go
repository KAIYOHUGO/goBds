package wss

import (
	"gobds/src/api"

	"github.com/gorilla/websocket"
)

var (
	ok   = api.Sstatus{Status: true}
	fail = api.Sstatus{Status: false}
)

func send(ws *websocket.Conn, v interface{}) {
	ws.WriteJSON(v)
}

// func fail(ws *websocket.Conn) {
// 	var v = api.Status{Status: false}
// 	send(ws, v)
// }

// func ok(ws *websocket.Conn) {
// 	var v = api.Status{Status: true}
// 	send(ws, v)
// }
