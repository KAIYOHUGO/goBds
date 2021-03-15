package wss

import (
	"gobds/src/usefull"
	"net/http"
)

func echoPlugin(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		usefull.Err("wss echo", err)
		return
	}
	ws.Close()

}
