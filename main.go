package main

import (
	"gobds/src/gc"
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
	defer func() {
		if err := recover(); err != nil {
			utils.Err("exit at err", err)
			os.Exit(10)
		}
		gc.GC()
		utils.Log("exit")
	}()
	select {}
}
