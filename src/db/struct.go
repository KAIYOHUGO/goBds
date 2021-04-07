package db

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"gobds/src/hoster"
	"gobds/vendor/github.com/syndtr/goleveldb/leveldb"
	"os"
)

// database struct
type DB struct {
	dblist map[string]*leveldb.DB
}
type User struct {
	Name          string
	Password      []byte
	OwnServerList []string
	Permission    int8
}

func (v *DB) init() {
	dbAccout, err := leveldb.OpenFile("./db/accout.db", nil)
	if err != nil {
		os.Exit(15)
	}
	dbServer, err := leveldb.OpenFile("./db/server.db", nil)
	if err != nil {
		os.Exit(15)
	}
	defer func() {
		fmt.Println("close")
	}()
	if ok, _ := dbAccout.Has([]byte("admire"), nil); !ok {
		sha.Reset()
		sha.Write([]byte("12345678"))
		dbAccout.Put([]byte("admire"), sha.Sum(nil), nil)
	}
	if ok, _ := dbServer.Has([]byte("TOL"), nil); !ok {
		encoder.Encode(hoster.List{
			Path: "./testserver.sh",
		})
		dbServer.Put([]byte("TOL"), encode.Bytes(), nil)
	}
	it := dbServer.NewIterator(nil, nil)
	fmt.Println("start it")
	for it.Next() {
		var decode hoster.List
		gob.NewDecoder(bytes.NewBuffer(it.Value())).Decode(&decode)
		hoster.ServerList[string(it.Key())] = &decode
	}
	it.Release()
	v.dblist["accout"] = dbAccout
	dbServer.Close()
}
func (v *DB) Search(d string, k []byte) ([]byte, error) {
	db, ok := v.dblist[d]
	if !ok {
		return []byte{}, errors.New("not find database")
	}
	b, err := db.Get(k, nil)
	if err != nil {
		return []byte{}, errors.New("not find in database")
	}
	return b, nil

}

// func (v *DB) init() {

// }
// func (v *DB) init() {

// }
