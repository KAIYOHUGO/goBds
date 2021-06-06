package api

import (
	"gobds/src/config"
	"gobds/src/database"
	"gobds/src/utils"
	"net/http"
)

func POSTSession(j *Request, w http.ResponseWriter, r *http.Request) (*API, error) {
	name, password, ok := r.BasicAuth()
	if !ok {
		return StatusBadRequest, nil
	}
	if !utils.IsExist(name, password) {
		return StatusUnprocessableEntity, nil
	}
	var a config.Account
	if err := database.Read(database.DB["account"], name, &a); err != nil {
		return StatusUnauthorized, err
	}
	if utils.Sha1([]byte(password)) != a.Password {
		return StatusUnauthorized, nil
	}
	session, err := database.NewSession(a)
	if err != nil {
		return StatusInternalServerError, err
	}
	w.Header().Set("Authorization", "Bearer "+session)
	return &API{
		Status: http.StatusCreated,
		Body: &Response{
			Messenge: "login suceess",
		},
	}, nil
}
func DELETESession(j *Request, w http.ResponseWriter, r *http.Request) (*API, error) {
	_, session, ok := GetLoginSession(r)
	if !ok {
		return StatusUnauthorized, nil
	}
	if err := database.DelSession(session); err != nil {
		return StatusInternalServerError, err
	}
	return &API{
		Status: http.StatusOK,
		Body:   &Response{Messenge: "Logout suceess"},
	}, nil
}
