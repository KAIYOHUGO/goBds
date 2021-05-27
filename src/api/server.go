package api

import (
	"gobds/src/console"
	"gobds/src/database"
	"gobds/src/utils"
	"net/http"
)

func PostServer(j *Request, w http.ResponseWriter, r *http.Request) *API {
	if _, _, ok := GetLoginSession(r); !ok {
		return StatusUnauthorized
	}
	if !utils.IsExist(j.Server) {
		return StatusUnprocessableEntity
	}
	if database.Has(database.DB["server"], j.Server) {
		return StatusConflict
	}
	wrapper := console.NewWrapper("")
	if database.Write(database.DB["server"], j.Server, wrapper) != nil {
		return StatusInternalServerError
	}
	console.ServerList[j.Server] = wrapper
	return &API{
		Status: http.StatusOK,
		Body: &Response{
			Messenge: "Server create suceess",
		},
	}
}
