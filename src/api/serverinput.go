package api

import (
	"encoding/base64"
	"gobds/src/config"
	"gobds/src/console"
	"gobds/src/database"
	"gobds/src/utils"
	"net/http"

	"github.com/dgraph-io/badger/v3"
	"github.com/gorilla/mux"
)

func POSTServerInput(j *Request, w http.ResponseWriter, r *http.Request) (*API, error) {
	if _, _, ok := GetLoginSession(r); !ok {
		return StatusUnauthorized, nil
	}
	vars := mux.Vars(r)
	if !utils.IsExist(vars["ServerID"]) {
		return StatusUnprocessableEntity, nil
	}
	serverid, err := base64.URLEncoding.DecodeString(vars["ServerID"])
	if err != nil {
		return StatusUnprocessableEntity, err
	}
	var server config.Server
	if err := database.Read(database.DB["server"], string(serverid), &server); err != nil {
		if err == badger.ErrInvalidKey {
			return StatusNotFound, err
		}
		return StatusInternalServerError, err
	}
	if !console.ServerList[string(serverid)].InputQueue(j.Input) {
		return StatusServiceUnavailable, nil
	}
	return &API{
		Status: http.StatusAccepted,
		Body: &Response{
			Messenge: "Input push to input queue suceess",
		},
	}, nil
}
