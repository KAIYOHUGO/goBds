package api

import (
	"errors"
	"gobds/src/config"
	"gobds/src/database"
	"gobds/src/utils"
	"net/http"
)

type Request struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Server   string `json:"server,omitempty"`
}

func GetToken(v string) (string, error) {
	if !utils.IsExist(v) {
		return "", errors.New("not exist")
	}
	if len(v) < 7 || v[:6] != "Bearer" {
		return "", errors.New("not bearer")
	}
	return v[7:], nil
}

func GetLoginSession(r *http.Request) (config.Session, string, bool) {
	token, err := GetToken(r.Header.Get("Authorization"))
	if err != nil {
		return config.Session{}, "", false
	}
	if !utils.IsExist(token) {
		return config.Session{}, "", false
	}
	session, err := database.GetSession(token)
	if err != nil {
		return config.Session{}, "", false
	}
	return session, token, true
}
