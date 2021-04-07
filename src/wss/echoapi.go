package wss

import (
	"gobds/src/hoster"
	"gobds/src/usefull"
	"net/http"

	"github.com/gorilla/websocket"
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
	for {
		var (
			client Client
		)
		if err = ws.ReadJSON(&client); err != nil {
			if websocket.IsCloseError(err) {
				return
			}
			ws.WriteJSON(fail)
		}
		user, err := Session.Get(client.Session)
		if err != nil {
			switch client.Type {
			case "login":
				if client.Password == "" {
					ws.WriteJSON(unFind)
					continue
				}

			default:
				ws.WriteJSON(noPermission)
			}
			continue
		}
		switch client.Type {
		case "event":
			if client.Event == "" || client.ServerName == "" {
				ws.WriteJSON(unFind)
				continue
			}
			for _, el := range user.info.OwnServerList {
				if client.ServerName == el {
					hoster.ServerList[el].EventChan <- client.Event
					break
				}
			}

		default:
			ws.WriteJSON(unFind)
		}
	}
}
