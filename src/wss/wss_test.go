package wss

import (
	"fmt"
	"gobds/src/config"
	"gobds/src/console"
	"gobds/src/database"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func TestServerTerminal(t *testing.T) {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.HandleFunc("/wss/server/{ServerID}/terminal/{SessionID}", ServerTerminal)
	s := httptest.NewServer(r)
	defer s.Close()
	d, err := database.NewSession(config.Account{})
	if err != nil {
		t.Fatal(err)
	}
	token := database.NewToken(config.ServerIDLen)
	database.Write(database.DB["server"], token, config.Server{Name: "test"})
	console.ServerList["test"] = console.NewWrapper(config.TestServerFile)
	defer console.ServerList["test"].Console.Kill()
	ws, b, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws%s/wss/server/%s/terminal/%s", strings.TrimPrefix(s.URL, "http"), token, d), nil)
	if err != nil {
		t.Fatal(err, b)
	}
	console.ServerList["test"].InputQueue("$start")
	console.ServerList["test"].InputQueue("stop")
	for i := 0; i < 15; i++ {
		mt, p, err := ws.ReadMessage()
		t.Log(mt)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(string(p))
	}
}
