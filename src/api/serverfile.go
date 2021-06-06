package api

import (
	"encoding/base64"
	"gobds/src/config"
	"gobds/src/database"
	"gobds/src/fileserver"
	"gobds/src/utils"
	"net/http"

	"github.com/dgraph-io/badger/v3"
	"github.com/gorilla/mux"
)

// read file, need Path
func GETServerFile(_ *Request, w http.ResponseWriter, r *http.Request) (*API, error) {
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
	if err := fileserver.NewFileServer(server.Path, vars["Path"], fileserver.ModeRead).Serve(w, r); err != nil {
		return StatusInternalServerError, err
	}
	return nil, nil
}

// create file, need Path & Type
func POSTServerFile(_ *Request, w http.ResponseWriter, r *http.Request) (*API, error) {
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
	var mode int8
	switch vars["Type"] {
	case "file":
		mode = fileserver.ModeCreateFile
	case "dir":
		mode = fileserver.ModeCreateDir
	default:
		return StatusUnprocessableEntity, nil
	}
	if err := fileserver.NewFileServer(server.Path, vars["Path"], mode).Serve(w, r); err != nil {
		return StatusInternalServerError, err
	}
	return nil, nil
}

// edit file, need Path
func PUTServerFile(_ *Request, w http.ResponseWriter, r *http.Request) (*API, error) {
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
	if err := fileserver.NewFileServer(server.Path, vars["Path"], fileserver.ModeWrite).Serve(w, r); err != nil {
		return StatusInternalServerError, err
	}
	return nil, nil
}

// delete file, need Path
func DELETEServerFile(_ *Request, w http.ResponseWriter, r *http.Request) (*API, error) {
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
	if err := fileserver.NewFileServer(server.Path, vars["Path"], fileserver.ModeDelete).Serve(w, r); err != nil {
		return StatusInternalServerError, err
	}
	return nil, nil
}

// rename file, need Path & NewPath
func PATCHServerFile(_ *Request, w http.ResponseWriter, r *http.Request) (*API, error) {
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
	fs := fileserver.NewFileServer(server.Path, vars["Path"], fileserver.ModeRename)
	fs.NewPath = vars["NewPath"]
	if err := fs.Serve(w, r); err != nil {
		return StatusInternalServerError, err
	}
	return nil, nil
}
