package api

import (
	"gobds/src/console"
	"net/http"
)

type Response struct {
	Error    string                      `json:"error,omitempty"`
	Messenge string                      `json:"messenge,omitempty"`
	Servers  map[string]*console.Wrapper `json:"servers,omitempty"`
}

type API struct {
	Status int
	Body   *Response
}

const (
	jsontype string = "application/json"
)

var (
	// http 400
	StatusBadRequest = &API{
		Status: http.StatusBadRequest,
		Body: &Response{
			Error:    "Bad request",
			Messenge: "Server can't decode request",
		},
	}
	// http 401
	StatusUnauthorized = &API{
		Status: http.StatusUnauthorized,
		Body: &Response{
			Error:    "Unauthorized",
			Messenge: "Login fail,check password or username again",
		},
	}
	StatusNotFound = &API{
		Status: http.StatusNotFound,
		Body: &Response{
			Error:    "Not Found",
			Messenge: "Resource not found in the database",
		},
	}
	// http 409
	StatusConflict = &API{
		Status: http.StatusConflict,
		Body: &Response{
			Error:    "Conflict",
			Messenge: "Resource already exist",
		},
	}
	// http 422
	StatusUnprocessableEntity = &API{
		Status: http.StatusUnprocessableEntity,
		Body: &Response{
			Error:    "Unprocessable Entity",
			Messenge: "Some filed are missing",
		},
	}
	// http 500
	StatusInternalServerError = &API{
		Status: http.StatusInternalServerError,
		Body: &Response{
			Error:    "Internal Server Error",
			Messenge: "Unknow server error",
		},
	}
)
