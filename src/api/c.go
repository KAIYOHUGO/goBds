package api

// Client ...
// type =password,cmd,event
// if type=server,need server,cmd
type Client struct {
	Type     string
	Name     string
	Password string
	Cmd      string
}
