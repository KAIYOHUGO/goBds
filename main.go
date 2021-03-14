package main

import (
	"fmt"
	"gobds/src/db"
	"gobds/src/hoster"
	"gobds/src/usefull"
	"gobds/src/wss"
	"os"
)

var testmain = make(chan struct{})

func main() {
	usefull.Log("starting ...")
	db.Run()
	usefull.Log("start server")
	hoster.Run()
	usefull.Log("start wss")
	go wss.Run()
	testmain <- struct{}{}
	for {
		var n string
		fmt.Scanln(&n)
		fmt.Println(n)
	}
	usefull.Wan("unknow error")
	os.Exit(3)
}
