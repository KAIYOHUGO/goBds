package db

import (
	_ "crypto/sha1"
)

var (
	err      error
	DataBase DB
)

func Run() {
	DataBase = DB{}
	DataBase.init()
}
func GC() {
	for _, v := range DataBase.dblist {
		v.Close()
	}
}
