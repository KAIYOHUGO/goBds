package db

import (
	"encoding/gob"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	err error
	db  *leveldb.DB
)

func Run() {
	db, err := leveldb.OpenFile("/db/accout.db", nil)
	if err != nil {
		return
	}
	defer db.Close()
	if ok, _ := db.Has([]byte("admire"), nil); !ok {
		db.Put([]byte("admire"), []byte(""), nil)
	}
	gob.NewEncoder(nil)
	// db.Put([]byte("ad"))
}
