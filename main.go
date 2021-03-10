package main

import (
	"fmt"
	"gobds/src/hoster"
	"gobds/src/msg"
	"gobds/src/wss"
	"os"
)

func main() {
	msg.Log("starting ...")
	msg.Log("start server")
	hoster.Run()
	msg.Log("start wss")
	go wss.Run()
	for {
		var n string
		fmt.Scanln(&n)
		fmt.Println(n)
	}
	msg.Wan("unknow error")
	os.Exit(3)
}
