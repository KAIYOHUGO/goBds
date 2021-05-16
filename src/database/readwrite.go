package database

import (
	"github.com/dgraph-io/badger/v3"
)

// s must be a pointer type
func Read(d *badger.DB, k string, s interface{}) error {
	return d.View(func(txn *badger.Txn) error {
		t, err := txn.Get([]byte(k))
		if err != nil {
			return err
		}
		b, err := t.ValueCopy(nil)
		if err != nil {
			return err
		}
		err = Decode(b, s)
		if err != nil {
			return err
		}
		return nil
	})
}

func Write(d *badger.DB, k string, s interface{}) error {
	return d.Update(func(txn *badger.Txn) error {
		b, err := Encode(s)
		if err != nil {
			return err
		}
		return txn.Set([]byte(k), b)
	})
}
