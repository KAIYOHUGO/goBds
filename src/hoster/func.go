package hoster

import (
	"gobds/src/msg"
)

func listener() {
	go func() {
		for {
			select {
			case e := <-Event:
				el := ServerList[e.Name]
				msg.Log("chan event" + e.Cmd)
				switch e.Cmd {
				case "start":
					msg.Log("run start")
					go el.Start()
				case "stop":
					msg.Log("run stop")
					el.Stop()
				case "kill":
					msg.Log("run kill")
					el.Kill()
				default:
					msg.Err("unknow type", nil)
				}
			case e := <-Cmd:
				el := ServerList[e.Name]
				msg.Log("chan cmd")
				el.Cmd(e.Cmd)
			}
		}
	}()

}
