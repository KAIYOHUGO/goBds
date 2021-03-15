package db

import (
	"bytes"
	"crypto"
	"encoding/gob"
	"fmt"
	"gobds/src/hoster"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

var (
	err     error
	encode  bytes.Buffer
	encoder = gob.NewEncoder(&encode)
	sha     = crypto.SHA1.New()
)

func Run() {
	dbAccout, err := leveldb.OpenFile("./db/accout.db", nil)
	if err != nil {
		os.Exit(15)
	}
	dbServer, err := leveldb.OpenFile("./db/server.db", nil)
	if err != nil {
		os.Exit(15)
	}
	defer func() {
		dbAccout.Close()
		dbServer.Close()
		fmt.Println("close")
	}()
	if ok, _ := dbAccout.Has([]byte("admire"), nil); !ok {
		sha.Reset()
		sha.Write([]byte("12345678"))
		dbAccout.Put([]byte("admire"), sha.Sum(nil), nil)
	}
	if ok, _ := dbServer.Has([]byte("TOL"), nil); !ok {
		encoder.Encode(hoster.List{
			Path: "C:\\Users\\kymcm\\Documents\\VSCode\\gobds\\bds\\bedrock_server.exe",
		})
		dbServer.Put([]byte("TOL"), encode.Bytes(), nil)
	}
	it := dbServer.NewIterator(nil, nil)
	fmt.Println("start it")
	for it.Next() {
		var decode hoster.List
		gob.NewDecoder(bytes.NewBuffer(it.Value())).Decode(&decode)
		hoster.ServerList[string(it.Key())] = &decode
		fmt.Println(string(it.Key()))
	}
	it.Release()
}
