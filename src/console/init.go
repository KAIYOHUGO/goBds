package console

import (
	"gobds/src/config"
	"gobds/src/database"

	"github.com/dgraph-io/badger/v3"
)

var ServerList map[string]*Wrapper

// init hoster server list
func init() {
	ServerList = make(map[string]*Wrapper)
	database.DB["server"].View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			v, err := it.Item().ValueCopy(nil)
			if err != nil {
				return err
			}
			var s config.Server
			err = database.Decode(v, &s)
			if err != nil {
				return err
			}
			ServerList[string(it.Item().KeyCopy(nil))] = NewWrapper(s.Path + config.DefaultStartScriptName)
		}
		return nil
	})

}
