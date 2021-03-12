package main

import (
	"fmt"
	"gobds/src/hoster"
	"gobds/src/msg"
	"gobds/src/wss"
	"os"
)

var testmain = make(chan struct{})

func main() {
	msg.Log("starting ...")
	msg.Log("start server")
	hoster.Run()
	msg.Log("start wss")
	go wss.Run()
	testmain <- struct{}{}
	for {
		var n string
		fmt.Scanln(&n)
		fmt.Println(n)
	}
	msg.Wan("unknow error")
	os.Exit(3)
}
