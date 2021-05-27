package api

import (
	"gobds/src/config"
	"gobds/src/database"
	"gobds/src/utils"
	"net/http"
)

func POSTSession(j *Request, w http.ResponseWriter, r *http.Request) *API {
	name, password, ok := r.BasicAuth()
	if !ok {
		panic(StatusBadRequest)
	}
	if !utils.IsExist(name, password) {
		panic(StatusUnprocessableEntity)
	}
	var a config.Account
	if err := database.Read(database.DB["account"], name, &a); err != nil {
		return StatusUnauthorized
	}
	if utils.Sha1([]byte(password)) != a.Password {
		return StatusUnauthorized
	}
	session, err := database.NewSession(a)
	if err != nil {
		return StatusInternalServerError
	}
	w.Header().Set("Authorization", "Bearer "+session)
	return &API{
		Status: http.StatusCreated,
		Body: &Response{
			Messenge: "login suceess",
		},
	}
}
func DELETESession(j *Request, w http.ResponseWriter, r *http.Request) *API {
	_, session, ok := GetLoginSession(r)
	if !ok {
		return StatusUnauthorized
	}
	if err := database.DelSession(session); err != nil {
		return StatusInternalServerError
	}
	return &API{
		Status: http.StatusOK,
		Body:   &Response{Messenge: "Logout suceess"},
	}
}
