package wss

import (
	"gobds/src/config"
	"gobds/src/database"
	"gobds/src/utils"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var (
	err     error
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
	server []*config.Server
)

func init() {
	database.DB["ServerID"].Read(&server)
}

// Run ...
// start wss server
func ServerTerminal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ws, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		utils.Err("can not start ws", err)
		return
	}
	defer ws.Close()
	for _, v := range server {
		if v.Name == vars["ServerId"] {
			break
		}
	}
}
