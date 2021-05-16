package wss

import (
	"gobds/src/config"
	"gobds/src/database"
	"gobds/src/hoster"
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
			ws.WriteControl(websocket.CloseMessage, []byte(err.(error).Error()), time.Now().Add(config.WSHandshakeTimeout))
		}
		ws.Close()
	}()
	// updata session ttl
	_, err = database.GetSession(vars["SessionID"])
	if err != nil {
		panic(err)
	}
	// get server info
	var s hoster.List
	err = database.Read(database.DB["server"], vars["ServerID"], s)
	if err != nil {
		panic(err)
	}
	m, l := s.Broadcast.New()
	defer s.Broadcast.Close(l)
	for v := range m {
		if _, _, err := ws.NextReader(); err != nil {
			return
		}
		ws.WriteMessage(websocket.TextMessage, []byte(v.(string)))
	}
}
