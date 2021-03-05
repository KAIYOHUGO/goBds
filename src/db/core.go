package db

import (
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	err error
)

func Run() {
	s, err := leveldb.OpenFile("/db/accout.db", nil)
	if err != nil {
		return
	}
	defer s.Close()

}
