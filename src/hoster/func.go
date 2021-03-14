package hoster

import "gobds/src/usefull"

func listener() {
	go func() {
		for {
			select {
			case e := <-Event:
				el := ServerList[e.Name]
				usefull.Log("chan event" + e.Cmd)
				switch e.Cmd {
				case "start":
					usefull.Log("run start")
					go el.Start()
				case "stop":
					usefull.Log("run stop")
					el.Stop()
				case "kill":
					usefull.Log("run kill")
					el.Kill()
				default:
					usefull.Err("unknow type", nil)
				}
			case e := <-Cmd:
				el := ServerList[e.Name]
				usefull.Log("chan cmd")
				el.Cmd(e.Cmd)
			}
		}
	}()

}
