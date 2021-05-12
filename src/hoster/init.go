package hoster

import (
	"bytes"
	"encoding/gob"
	"gobds/src/database"

	"github.com/dgraph-io/badger/v3"
)

// init hoster server list
func init() {
	database.DB["server"].View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			v, err := it.Item().ValueCopy(nil)
			if err != nil {
				return err
			}
			var s List
			err = gob.NewDecoder(bytes.NewBuffer(v)).Decode(&s)
			if err != nil {
				return err
			}
			ServerList[string(it.Item().KeyCopy(nil))] = &s

		}
		return nil
	})
}
