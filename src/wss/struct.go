package wss

var (
	ok           = Server{Code: 200, Messenge: "ok"}
	fail         = Server{Code: 500, Messenge: "server fail"}
	noPermission = Server{Code: 400, Messenge: "you don't have permission"}
	unFind       = Server{Code: 404, Messenge: "unfind json node"}
)

// Client ...
// type =login,cmd,event
// if type=server,need server,cmd
type Client struct {
	Type     string `json:"type,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Cmd      string `json:"cmd,omitempty"`
}

type Server struct {
	// `json:",omitempty"`
	// like http code
	Code     int    `json:"code,omitempty"`
	Messenge string `json:"messenge,omitempty"`
	// terminal,login,
	// Type   string `json:"type,omitempty"`
	Session    string   `json:"session,omitempty"`
	Terminal   []string `json:"terminal,omitempty"`
	ServerList []string `json:"serverlist,omitempty"`
}

// func send(ws *websocket.Conn, v interface{}) {
// 	ws.WriteJSON(v)
// }

// func fail(ws *websocket.Conn) {
// 	var v = api.Status{Status: false}
// 	send(ws, v)
// }

// func ok(ws *websocket.Conn) {
// 	var v = api.Status{Status: true}
// 	send(ws, v)
// }

// Sstatus ...
// return bool,true mean something is right
type Sstatus struct {
	Status bool
}
