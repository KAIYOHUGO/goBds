package db

import (
	"bytes"
	"crypto"
	_ "crypto/sha1"
	"encoding/gob"
	"errors"
	"fmt"
	"gobds/src/hoster"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
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
	var (
		encode bytes.Buffer
		sha    = crypto.SHA1.New()
	)
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
		dbServer.Close()
	}()
	if ok, _ := dbAccout.Has([]byte("admire"), nil); !ok {
		sha.Reset()
		sha.Write([]byte("12345678"))
		dbAccout.Put([]byte("admire"), sha.Sum(nil), nil)
	}
	if ok, _ := dbServer.Has([]byte("TOL"), nil); !ok {
		gob.NewEncoder(&encode).Encode(hoster.List{
			Path: "./testserver.sh",
		})
		dbServer.Put([]byte("TOL"), encode.Bytes(), nil)
	}
	it := dbServer.NewIterator(nil, nil)
	defer it.Release()
	fmt.Println("start it")
	for it.Next() {
		var decode hoster.List
		gob.NewDecoder(bytes.NewBuffer(it.Value())).Decode(&decode)
		hoster.ServerList[string(it.Key())] = &decode
	}
	v.dblist["accout"] = dbAccout
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
func (v *DB) SearchStruct(d string, k []byte, s *struct{}) error {
	db, ok := v.dblist[d]
	if !ok {
		return errors.New("not find database")
	}
	b, err := db.Get(k, nil)
	if err != nil {
		return errors.New("not find in database")
	}
	return gob.NewDecoder(bytes.NewBuffer(b)).Decode(s)
}

// func (v *DB) init() {

// }
// func (v *DB) init() {

// }
