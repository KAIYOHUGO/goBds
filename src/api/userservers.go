package api

import (
	"encoding/base64"
	"gobds/src/console"
	"gobds/src/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GETUserServers(j *Request, w http.ResponseWriter, r *http.Request) *API {
	vars := mux.Vars(r)
	user, _, ok := GetLoginSession(r)
	if !ok {
		return StatusUnauthorized
	}
	if !utils.IsExist(vars["UserID"]) {
		return StatusBadRequest
	}
	userid, err := base64.URLEncoding.DecodeString(vars["UserID"])
	if err != nil {
		return StatusBadRequest
	}
	if user.Name != string(userid) {
		return StatusNotFound
	}
	return &API{
		Status: http.StatusOK,
		Body: &Response{
			Messenge: "get servers success",
			Servers:  console.ServerList,
		},
	}
}
