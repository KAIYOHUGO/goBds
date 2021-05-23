package api

import (
	"gobds/src/config"
	"gobds/src/database"
	"gobds/src/utils"
	"net/http"
)

func POSTSession(w http.ResponseWriter, r *http.Request) {
	defer SendResponse(w, r)
	name, password, ok := r.BasicAuth()
	if !ok {
		panic(StatusBadRequest)
	}
	if !utils.IsExist(name, password) {
		panic(StatusUnprocessableEntity)
	}
	var a config.Account
	if err := database.Read(database.DB["account"], name, &a); err != nil {
		panic(StatusUnauthorized)
	}
	if utils.Sha1([]byte(password)) != a.Password {
		panic(StatusUnauthorized)
	}
	session, err := database.NewSession(a)
	if err != nil {
		panic(StatusInternalServerError)
	}
	w.Header().Set("Authorization", "Bearer "+session)
	panic(&API{
		Status: http.StatusCreated,
		ErrorMessenge: &Response{
			Messenge: "login suceess",
		},
	})
}
func DELETESession(w http.ResponseWriter, r *http.Request) {
	defer SendResponse(w, r)
	session, ok := GetLoginSession(r)
	if !ok {
		panic(StatusUnauthorized)
	}
	if err := database.DelSession(session); err != nil {
		panic(StatusUnauthorized)
	}
	panic(&API{
		Status:        http.StatusOK,
		ErrorMessenge: &Response{Messenge: "Logout suceess"},
	})
}
