package main

import (
	"fmt"
	"gobds/src/hoster"
	"gobds/src/utils"
	"math/rand"
	"os"
	"time"
)

func main() {
	utils.Log("starting ...")
	defer func() {
	}()
	rand.Seed(time.Now().UnixNano())
	utils.Log("start server")
	hoster.Run()
	utils.Log("start wss")
	for {
		var n string
		fmt.Scanln(&n)
		fmt.Println(n)
	}
	utils.Wan("unknow error")
	os.Exit(3)
}
