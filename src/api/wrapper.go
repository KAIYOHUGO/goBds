package api

import (
	"encoding/json"
	"gobds/src/utils"
	"io"
	"net/http"
)

func Wrapper(fn func(j *Request, w http.ResponseWriter, r *http.Request) *API) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", jsontype)
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(StatusInternalServerError.Status)
				json.NewEncoder(w).Encode(StatusInternalServerError.Body)
			}
		}()
		var j Request
		if err := json.NewDecoder(r.Body).Decode(&j); err != nil && err != io.EOF {
			w.WriteHeader(StatusBadRequest.Status)
			json.NewEncoder(w).Encode(StatusBadRequest.Body)
			utils.Err("api wrapper", err)
			return
		}
		s := fn(&j, w, r)
		w.WriteHeader(s.Status)
		if err := json.NewEncoder(w).Encode(s.Body); err != nil {
			w.WriteHeader(StatusInternalServerError.Status)
			json.NewEncoder(w).Encode(StatusInternalServerError.Body)
			utils.Err("api wrapper", err)
		}
	}
}
