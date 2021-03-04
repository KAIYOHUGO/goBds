package wss

import (
	"gobds/src/api"
	"gobds/src/msg"
	"gobds/src/runner"
)

func get(s api.Gclient, login *bool) interface{} {
	msg.Log("get !")
	if *login {
		switch s.Type {
		case "cmd":
			msg.Log("cmd:" + s.Cmd)
			runner.Cmd <- api.Prun{
				Name: s.Name,
				Cmd:  s.Cmd,
			}
			return ok
		case "event":
			msg.Log("term:" + s.Cmd)
			runner.Event <- api.Prun{
				Name: s.Name,
				Cmd:  s.Cmd,
			}
			return ok
		case "server":
			return api.Sinfo{
				Status: true,
				Info:   runner.List[s.Name].Status,
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
