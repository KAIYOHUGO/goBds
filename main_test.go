package main

import (
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestMain(t *testing.T) {
	go main()
	time.Sleep(time.Second * 20)

}

func TestRouter(t *testing.T) {
	go router()
	ws, _, err := websocket.DefaultDialer.Dial("ws://127.0.0.1", nil)
	if err != nil {
		return
	}
	_, p, err := ws.ReadMessage()
	if err != nil {
		return
	}
	println(p)
}
