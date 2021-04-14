package wss

var (
	ok           = Server{Code: 200, Messenge: "ok"}
	badRequest   = Server{Code: 400, Messenge: "bad Request! can you follow api?"}
	noPermission = Server{Code: 401, Messenge: "you don't have permission"}
	notFind      = Server{Code: 404, Messenge: "can not find json node or value"}
	tooLarge     = Server{Code: 413, Messenge: "payload too large"}
	fail         = Server{Code: 500, Messenge: "server fail"}
)

// Client ...
// if type=server,need server,cmd
type Client struct {
	// type =login,cmd,event
	Type       string `json:"type,omitempty"`
	Session    string `json:"session,omitempty"`
	Name       string `json:"name,omitempty"`
	Password   string `json:"password,omitempty"`
	ServerName string `json:"servername,omitempty"`
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
