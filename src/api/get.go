package api

// Gclient ...
// type =password,cmd,event
// if type=server,need server,cmd
type Gclient struct {
	Type     string
	Name     string
	Password string
	Cmd      string
}
