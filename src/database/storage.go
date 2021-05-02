package database

import "gobds/src/storage"

var DB = make(map[string]*storage.Storage)

func init() {
	DB["ServerId"], _ = storage.Open("./server.json")
	DB["Account"], _ = storage.Open("./account.json")
}
