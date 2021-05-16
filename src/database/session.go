package database

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"gobds/src/config"
	"math/rand"
	"time"

	"github.com/dgraph-io/badger/v3"
)

// get session,input session id,return (struct,error)
func GetSession(v string) (config.Session, error) {
	var s config.Session
	err := DB["session"].Update(func(txn *badger.Txn) error {
		t, err := txn.Get([]byte(v))
		if err != nil {
			return err
		}
		v, err := t.ValueCopy(nil)
		if err != nil {
			return err
		}
		err = gob.NewDecoder(bytes.NewBuffer(v)).Decode(&s)
		if err != nil {
			return err
		}
		return txn.SetEntry(badger.NewEntry(t.KeyCopy(nil), v).WithTTL(config.MaxSessionLiveTime))
	})
	return s, err
}

func NewSession(v config.Account) (string, error) {
	rand.Seed(time.Now().UnixNano())
	k := make([]byte, config.SessionIDLen)
	rand.Read(k)
	s := base64.URLEncoding.EncodeToString(k)
	b, err := Encode(config.Session{Name: v.Name})
	if err != nil {
		return "", err
	}
	return s, DB["session"].Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(s), b)
	})
}

// delete session ,input session id,return error
func DelSession(v string) error {
	return DB["session"].Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(v))
	})
}
