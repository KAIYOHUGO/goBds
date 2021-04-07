package db

import (
	"bytes"
	"crypto"
	_ "crypto/sha1"
	"encoding/gob"
)

var (
	err      error
	encode   bytes.Buffer
	encoder  = gob.NewEncoder(&encode)
	sha      = crypto.SHA1.New()
	DataBase DB
)

func Run() {
	DataBase = DB{}
	DataBase.init()
	defer func() {

	}()
}
