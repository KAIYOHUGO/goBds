package wss

import (
	"gobds/src/api"
	"gobds/src/hoster"
	"gobds/src/msg"
)

func get(s Client, login *bool) interface{} {
	msg.Log("get !")
	if *login {
		switch s.Type {
		case "cmd":
			msg.Log("cmd:" + s.Cmd)
			hoster.Cmd <- api.Prun{
				Name: s.Name,
				Cmd:  s.Cmd,
			}
			return ok
		case "event":
			msg.Log("term:" + s.Cmd)
			hoster.Event <- api.Prun{
				Name: s.Name,
				Cmd:  s.Cmd,
			}
			return ok
		case "server":
			return Sinfo{
				Status: true,
				Info:   hoster.ServerList[s.Name].Status,
			}
		default:
			return fail
		}
	} else {
		if s.Type == "password" && s.Password == "12345678" {
			*login = true
			return ok
		}
		return fail

	}
}
