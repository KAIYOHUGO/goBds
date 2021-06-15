package main

import (
	"testing"
	"time"
)

// func TestMain(t *testing.T) {
// 	go main()
// 	time.Sleep(time.Second * 20)

// }

func TestRouter(t *testing.T) {
	go main()
	time.Sleep(time.Second * 10)
}
