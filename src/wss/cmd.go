package wss

import (
	"gobds/src/usefull"
	"net/http"
)

func echoCmd(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		usefull.Err("wss echo", err)
		return
	}
	defer ws.Close()

}
