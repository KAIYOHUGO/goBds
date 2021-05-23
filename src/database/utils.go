package database

import "github.com/dgraph-io/badger/v3"

func Has(db *badger.DB, v string) bool {
	err := db.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(v))
		return err
	})
	return err != badger.ErrKeyNotFound
}

func Delete(db *badger.DB, v string) error {
	return db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte(v))
	})
}
