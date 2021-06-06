package api

import (
	"gobds/src/config"
	"gobds/src/console"
	"gobds/src/database"
	"gobds/src/utils"
	"net/http"
)

func PostServer(j *Request, w http.ResponseWriter, r *http.Request) (*API, error) {
	if _, _, ok := GetLoginSession(r); !ok {
		return StatusUnauthorized, nil
	}
	if !utils.IsExist(j.Server) {
		return StatusUnprocessableEntity, nil
	}
	if database.Has(database.DB["server"], j.Server) {
		return StatusConflict, nil
	}
	if err := database.Write(database.DB["server"], j.Server, &config.Server{
		Name:    j.Server,
		Path:    j.Path,
		Command: j.Command,
	}); err != nil {
		return StatusInternalServerError, err
	}
	console.ServerList[j.Server] = console.NewWrapper(j.Path, j.Command)
	return &API{
		Status: http.StatusCreated,
		Body: &Response{
			Messenge: "Server create suceess",
		},
	}, nil
}

func DeleteServer(j *Request, w http.ResponseWriter, r *http.Request) (*API, error) {
	if _, _, ok := GetLoginSession(r); !ok {
		return StatusUnauthorized, nil
	}
	if !utils.IsExist(j.Server) {
		return StatusUnprocessableEntity, nil
	}
	if !database.Has(database.DB["server"], j.Server) {
		return StatusNotFound, nil
	}
	if err := database.Delete(database.DB["server"], j.Server); err != nil {
		return StatusInternalServerError, err
	}
	server, ok := console.ServerList[j.Server]
	if !ok {
		return StatusInternalServerError, nil
	}
	server.GC()
	return &API{
		Status: http.StatusOK,
		Body: &Response{
			Messenge: "Server delete suceess",
		},
	}, nil
}
