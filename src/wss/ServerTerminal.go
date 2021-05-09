package wss

import (
	"bytes"
	"encoding/gob"
	"gobds/src/config"
	"gobds/src/database"
	"gobds/src/utils"
	"net/http"
	"time"

	"github.com/dgraph-io/badger/v3"
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
	err = database.DB["session"].Update(func(txn *badger.Txn) error {
		t, err := txn.Get([]byte(vars["SessionID"]))
		if err != nil {
			return err
		}
		v, err := t.ValueCopy(nil)
		if err != nil {
			return err
		}
		badger.NewEntry(t.KeyCopy(nil), v).WithTTL(config.MaxSessionLiveTime)
		return nil
	})
	if err != nil {
		panic(err)
	}
	// get server info
	var s database.Server
	err = database.DB["server"].View(func(txn *badger.Txn) error {
		t, err := txn.Get([]byte(vars["ServerID"]))
		if err != nil {
			return err
		}
		v, err := t.ValueCopy(nil)
		if err != nil {
			return err
		}
		err = gob.NewDecoder(bytes.NewBuffer(v)).Decode(&s)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	m := make(chan interface{})
	s.Broadcast.Add(m)
	for v := range m {
		if _, _, err := ws.NextReader(); err != nil {
			return
		}
		ws.WriteMessage(websocket.TextMessage, []byte(v.(string)))
	}
}
