package wss

import (
	"gobds/src/usefull"
	"net/http"
)

func echoAPI(w http.ResponseWriter, r *http.Request) {
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
	var (
		client Client
		server Server
		conn   bool = true
	)
	token := r.URL.Path[len("ws/api/"):]
	if user, err := Session.Get(token); err != nil {
		go func() {
			for conn {
				if err = ws.ReadJSON(&client); err != nil {
					conn = false
					break
				}
				if client.Type != "login" {
					ws.WriteJSON(noPermission)
					continue
				} else {
					if client.Password == "" {
						ws.WriteJSON(unFind)
						continue
					} else {
						client.Password
						continue
					}
				}
			}
		}()
	} else {

	}
	return
}
