package main

import (
	"gobds/src/gc"
	"gobds/src/hoster"
	"gobds/src/utils"
	"math/rand"
	"time"
)

func main() {
	utils.Log("starting ...")
	s := make(chan struct{})
	defer func() {
		close(s)
		gc.GC()
		utils.Log("exit")
	}()
	rand.Seed(time.Now().UnixNano())
	utils.Log("start hoster...")
	hoster.Run()
	utils.Log("start api server...")
	router()
	<-s
}
