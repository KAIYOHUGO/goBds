package wss

var (
	ok           = Server{Code: 200, Messenge: "ok"}
	fail         = Server{Code: 500, Messenge: "server fail"}
	noPermission = Server{Code: 400, Messenge: "you don't have permission"}
	unFind       = Server{Code: 404, Messenge: "unfind json node or value"}
)

// Client ...
// if type=server,need server,cmd
type Client struct {
	// type =login,cmd,event
	Type       string `json:"type,omitempty"`
	Session    string `json:"session,omitempty"`
	ServerName string `json:"servername,omitempty"`
	Password   string `json:"password,omitempty"`
	Cmd        string `json:"cmd,omitempty"`
	Event      string `json:"event,omitempty"`
	FileName   string `json:"filename,omitempty"`
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
