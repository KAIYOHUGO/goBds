package api

import (
	"encoding/base64"
	"gobds/src/config"
	"gobds/src/console"
	"gobds/src/database"
	"gobds/src/utils"
	"net/http"
	"os"
	"path/filepath"
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
		Command: j.Command,
	}); err != nil {
		return StatusInternalServerError, err
	}
	rootPath, err := filepath.Abs(config.ServerRootDir + base64.URLEncoding.EncodeToString([]byte(j.Server)))
	if err != nil {
		return StatusInternalServerError, err
	}
	if err := os.MkdirAll(rootPath, os.ModePerm); err != nil {
		return StatusInternalServerError, err
	}
	console.ServerList[j.Server] = console.NewWrapper(rootPath, j.Command)
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
	rootPath, _ := filepath.Abs(config.ServerRootDir + base64.URLEncoding.EncodeToString([]byte(j.Server)))
	if err := os.RemoveAll(rootPath); err != nil {
		return StatusInternalServerError, err
	}
	server.GC()
	return &API{
		Status: http.StatusOK,
		Body: &Response{
			Messenge: "Server delete suceess",
		},
	}, nil
}
