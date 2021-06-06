package api

import (
	"encoding/base64"
	"gobds/src/config"
	"gobds/src/database"
	"gobds/src/utils"
	"net/http"

	"github.com/dgraph-io/badger/v3"
	"github.com/gorilla/mux"
)

func GETUserServers(j *Request, w http.ResponseWriter, r *http.Request) (*API, error) {
	user, _, ok := GetLoginSession(r)
	if !ok {
		return StatusUnauthorized, nil
	}
	vars := mux.Vars(r)
	if !utils.IsExist(vars["UserID"]) {
		return StatusBadRequest, nil
	}
	userid, err := base64.URLEncoding.DecodeString(vars["UserID"])
	if err != nil {
		return StatusBadRequest, err
	}
	if user.Name != string(userid) {
		return StatusNotFound, nil
	}
	servers := make(map[string]*config.Server)
	if err := database.DB["server"].View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			b, err := it.Item().ValueCopy(nil)
			if err != nil {
				return err
			}
			var server config.Server
			if err := database.Decode(b, &server); err != nil {
				return err
			}
			servers[string(it.Item().KeyCopy(nil))] = &server
		}
		return nil
	}); err != nil {
		return StatusInternalServerError, err
	}
	return &API{
		Status: http.StatusOK,
		Body: &Response{
			Messenge: "get servers success",
			Servers:  servers,
		},
	}, nil
}
