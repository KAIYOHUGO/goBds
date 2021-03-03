package main

import (
	"gobds/src/msg"
	"gobds/src/runner"
	"gobds/src/wss"
	"os"
)

func main() {
	msg.Log("starting ...")
	msg.Log("start server")
	runner.Run()
	msg.Log("start wss")
	wss.Run()
	msg.Wan("unknow error")
	os.Exit(3)
}
