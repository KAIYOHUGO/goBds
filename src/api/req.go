package api

import (
	"errors"
	"gobds/src/database"
	"gobds/src/utils"
	"net/http"
)

func GetToken(v string) (string, error) {
	if !utils.IsExist(v) {
		return "", errors.New("not exist")
	}
	if len(v) < 7 || v[:6] != "Bearer" {
		return "", errors.New("not bearer")
	}
	return v[7:], nil
}

func GetLoginSession(r *http.Request) (string, bool) {
	token, err := GetToken(r.Header.Get("Authorization"))
	if err != nil {
		return "", false
	}
	if !utils.IsExist(token) {
		return "", false
	}
	if _, err := database.GetSession(token); err != nil {
		return "", false
	}
	return token, true
}
