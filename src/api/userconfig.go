package api

import (
	"gobds/src/config"
	"gobds/src/database"
	"gobds/src/utils"
	"net/http"
)

func GETUserConfig(w http.ResponseWriter, r *http.Request) {

}
func PUTUserConfig(j *Request, w http.ResponseWriter, r *http.Request) (*API, error) {
	session, token, ok := GetLoginSession(r)
	if !ok {
		return StatusUnauthorized, nil
	}
	if !utils.IsExist(j.Password, j.NewPassword) {
		return StatusUnprocessableEntity, nil
	}
	if utils.Sha1([]byte(j.Password)) != session.Password {
		return StatusUnauthorized, nil
	}
	if err := database.DelSession(token); err != nil {
		return StatusInternalServerError, err
	}
	database.Write(database.DB["account"], j.Name, config.Account{
		Name:     j.Name,
		Password: j.NewPassword,
	})
	newSession, err := database.NewSession(config.Account{
		Name:     j.Name,
		Password: j.NewPassword,
	})
	if err != nil {
		return StatusInternalServerError, err
	}
	w.Header().Set("Authorization", "Bearer "+newSession)
	return &API{
		Status: http.StatusCreated,
		Body: &Response{
			Messenge: "update suceess",
		},
	}, nil
}
