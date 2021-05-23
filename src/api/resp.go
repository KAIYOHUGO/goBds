package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Error    string `json:"error,omitempty"`
	Messenge string `json:"messenge,omitempty"`
}

type API struct {
	Status        int
	ErrorMessenge *Response
}

const (
	jsontype string = "application/json"
)

var (
	// http 400
	StatusBadRequest = &API{
		Status: http.StatusBadRequest,
		ErrorMessenge: &Response{
			Error:    "Bad request",
			Messenge: "Server can't decode request",
		},
	}
	// http 401
	StatusUnauthorized = &API{
		Status: http.StatusUnauthorized,
		ErrorMessenge: &Response{
			Error:    "Unauthorized",
			Messenge: "Login fail,check password or username again",
		},
	}
	StatusNotFound = &API{
		Status: http.StatusNotFound,
		ErrorMessenge: &Response{
			Error:    "Not Found",
			Messenge: "Resource not found in the database",
		},
	}
	// http 409
	StatusConflict = &API{
		Status: http.StatusConflict,
		ErrorMessenge: &Response{
			Error:    "Conflict",
			Messenge: "Resource already exist",
		},
	}
	// http 422
	StatusUnprocessableEntity = &API{
		Status: http.StatusUnprocessableEntity,
		ErrorMessenge: &Response{
			Error:    "Unprocessable Entity",
			Messenge: "Some filed are missing",
		},
	}
	// http 500
	StatusInternalServerError = &API{
		Status: http.StatusInternalServerError,
		ErrorMessenge: &Response{
			Error:    "Internal Server Error",
			Messenge: "Unknow server error",
		},
	}
)

func SendResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", jsontype)
	if err := recover(); err != nil {
		e, ok := err.(*API)
		if !ok {
			w.WriteHeader(StatusInternalServerError.Status)
			json.NewEncoder(w).Encode(StatusInternalServerError.ErrorMessenge)
			return
		}
		w.WriteHeader(e.Status)
		json.NewEncoder(w).Encode(e.ErrorMessenge)
	}
}
