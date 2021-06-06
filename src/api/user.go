package api

import (
	"gobds/src/config"
	"gobds/src/database"
	"gobds/src/utils"
	"net/http"
)

func POSTUser(j *Request, w http.ResponseWriter, r *http.Request) (*API, error) {
	if _, _, ok := GetLoginSession(r); !ok {
		return StatusUnauthorized, nil
	}
	if !utils.IsExist(j.Name, j.Password) {
		panic(StatusUnprocessableEntity)
	}
	if database.Has(database.DB["account"], j.Name) {
		return StatusConflict, nil
	}
	database.Write(database.DB["account"], j.Name, config.Account{Name: j.Name, Password: utils.Sha1([]byte(j.Password))})
	return &API{
		Status: http.StatusCreated,
		Body: &Response{
			Messenge: "User create suceess",
		},
	}, nil
}
func DELETEUser(j *Request, w http.ResponseWriter, r *http.Request) (*API, error) {
	if _, _, ok := GetLoginSession(r); !ok {
		return StatusUnauthorized, nil
	}
	if !utils.IsExist(j.Name) {
		return StatusUnprocessableEntity, nil
	}
	if !database.Has(database.DB["account"], j.Name) {
		return StatusNotFound, nil
	}
	if err := database.Delete(database.DB["account"], j.Name); err != nil {
		return StatusInternalServerError, err
	}
	return &API{
		Status: http.StatusOK,
		Body: &Response{
			Messenge: "Delete user success",
		},
	}, nil
}
