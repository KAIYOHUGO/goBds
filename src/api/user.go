package api

import (
	"gobds/src/config"
	"gobds/src/database"
	"gobds/src/utils"
	"net/http"
)

func POSTUser(j *Request, w http.ResponseWriter, r *http.Request) *API {
	if _, _, ok := GetLoginSession(r); !ok {
		panic(StatusUnauthorized)
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
		Body: &Response{
			Messenge: "User create suceess",
		},
	})
}
func DELETEUser(j *Request, w http.ResponseWriter, r *http.Request) *API {
	if _, _, ok := GetLoginSession(r); !ok {
		return StatusUnauthorized
	}
	if !utils.IsExist(j.Name) {
		return StatusUnprocessableEntity
	}
	if !database.Has(database.DB["account"], j.Name) {
		return StatusNotFound
	}
	if database.Delete(database.DB["account"], j.Name) != nil {
		return StatusInternalServerError
	}
	return &API{
		Status: http.StatusOK,
		Body: &Response{
			Messenge: "Delete user success",
		},
	}
}
