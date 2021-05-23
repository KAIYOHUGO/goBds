package api

import (
	"encoding/json"
	"gobds/src/config"
	"gobds/src/database"
	"gobds/src/utils"
	"net/http"
)

type ReqUser struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

func POSTUser(w http.ResponseWriter, r *http.Request) {
	defer SendResponse(w, r)
	if _, ok := GetLoginSession(r); !ok {
		panic(StatusUnauthorized)
	}
	var j ReqUser
	if json.NewDecoder(r.Body).Decode(&j) != nil {
		panic(StatusBadRequest)
	}
	if !utils.IsExist(j.Name, j.Password) {
		panic(StatusUnprocessableEntity)
	}
	if database.Has(database.DB["account"], j.Name) {
		panic(StatusConflict)
	}
	database.Write(database.DB["account"], j.Name, config.Account{Name: j.Name, Password: utils.Sha1([]byte(j.Password))})
	panic(&API{
		Status: http.StatusCreated,
		ErrorMessenge: &Response{
			Messenge: "User create suceess",
		},
	})
}
func DELETEUser(w http.ResponseWriter, r *http.Request) {
	defer SendResponse(w, r)
	if _, ok := GetLoginSession(r); !ok {
		panic(StatusUnauthorized)
	}
	var j ReqUser
	json.NewDecoder(r.Body).Decode(&j)
	if !utils.IsExist(j.Name) {
		panic(StatusUnprocessableEntity)
	}
	if !database.Has(database.DB["account"], j.Name) {
		panic(StatusNotFound)
	}
	if database.Delete(database.DB["account"], j.Name) != nil {
		panic(StatusInternalServerError)
	}
	panic(&API{
		Status: http.StatusOK,
		ErrorMessenge: &Response{
			Messenge: "Delete user success",
		},
	})
}
