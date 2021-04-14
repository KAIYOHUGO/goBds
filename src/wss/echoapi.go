package wss

import (
	"bytes"
	"crypto"
	_ "crypto/sha1"
	"gobds/src/config"
	"gobds/src/db"
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
				if client.Password == "" || client.Name == "" {
					ws.WriteJSON(badRequest)
					continue
				}
				if len(client.Password) > config.MaxAPIPayloadLen || len(client.Name) > config.MaxAPIPayloadLen {
					ws.WriteJSON(tooLarge)
					continue
				}
				pass, err := db.DataBase.Search("account", []byte(client.Name))
				if err != nil {
					ws.WriteJSON(notFind)
					continue
				}
				sha := crypto.SHA1.New()
				if sha.Write([]byte(client.Password)); !bytes.Equal(sha.Sum(nil), pass) {
					ws.WriteJSON(notFind)
					continue
				}
				s, err := Session.Add()
				if err != nil {
					ws.WriteJSON(fail)
					continue
				}
				ws.WriteJSON(Server{
					Code:       200,
					Messenge:   "login suceess",
					Session:    s,
					Terminal:   []string{},
					ServerList: []string{},
				})
				continue
			default:
				ws.WriteJSON(noPermission)
			}
			continue
		}
		switch client.Type {
		case "event":
			if client.Event == "" || client.ServerName == "" {
				ws.WriteJSON(badRequest)
				continue
			}
			for _, el := range user.info.OwnServerList {
				if client.ServerName == el {
					select {
					case hoster.ServerList[el].EventChan <- client.Event:
					default:
						ws.WriteJSON(fail)
					}
					break
				}
			}

		default:
			ws.WriteJSON(badRequest)
		}
	}
}
