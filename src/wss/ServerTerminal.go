package wss

import (
	"gobds/src/config"
	"gobds/src/console"
	"gobds/src/database"
	"gobds/src/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	upgrade = websocket.Upgrader{
		HandshakeTimeout: config.WSHandshakeTimeout,
		ReadBufferSize:   config.MaxWSBufferSize,
		WriteBufferSize:  config.MaxWSBufferSize,
		WriteBufferPool:  nil,
		Subprotocols:     []string{},
		Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
		},
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		EnableCompression: false,
	}
)

// Run ...
// start wss server
func ServerTerminal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ws, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		utils.Err("can not start ws", err)
		return
	}
	defer func() {
		if err := recover(); err != nil {
			ws.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(1000, err.(error).Error()), time.Now().Add(config.WSHandshakeTimeout))
		} else {
			ws.Close()
		}
	}()
	// updata session ttl
	_, err = database.GetSession(vars["SessionID"])
	if err != nil {
		panic(err)
	}
	// get server info
	var s config.Server
	err = database.Read(database.DB["server"], vars["ServerID"], &s)
	if err != nil {
		panic(err)
	}
	m, l := console.ServerList[s.Name].Broadcast.New()
	defer console.ServerList[s.Name].Broadcast.Close(l)
	for v := range m {
		if err := ws.WriteJSON(v.(console.Log)); err != nil {
			return
		}
	}
}
