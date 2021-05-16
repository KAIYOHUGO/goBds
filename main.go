package main

import (
	"gobds/src/gc"
	"gobds/src/hoster"
	"gobds/src/utils"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	utils.Log("starting ...")
	s := make(chan struct{})
	defer func() {
		if err := recover(); err != nil {
			utils.Err("exit at err", err)
			os.Exit(10)
		}
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
