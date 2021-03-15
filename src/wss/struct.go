package wss

var (
	ok           = Server{Status: 200}
	fail         = Server{Status: 500}
	noPermission = Server{Status: 400}
	unFind       = Server{Status: 404}
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
	Status int `json:"status,omitempty"`
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

// Sinfo ...
type Sinfo struct {
	Status bool
	Info   int8
}

// Slist ...
// return name + status array
type Slist []struct {
	Name   string
	Status bool
}

// Sterminal ...
// return line
type Sterminal struct {
	line string
}
