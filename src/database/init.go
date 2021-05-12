package database

import (
	"github.com/dgraph-io/badger/v3"
)

var DB map[string]*badger.DB

// init database
func init() {
	var err error
	dblist := []string{"server", "account"}
	DB = make(map[string]*badger.DB)
	for _, v := range dblist {
		DB[v], err = badger.Open(badger.DefaultOptions("./db/" + v))
		if err != nil {
			panic("can not open DB \n" + err.Error())
		}
	}
	DB["session"], err = badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		panic("can not start session DB" + err.Error())
	}
}
