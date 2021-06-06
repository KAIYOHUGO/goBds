package api

import (
	"encoding/json"
	"fmt"
	"gobds/src/utils"
	"io"
	"net/http"
	"runtime/debug"
)

// auto decode request to json
// if *API == nil will omit *API
// if error != nil will log error
func Wrapper(fn func(j *Request, w http.ResponseWriter, r *http.Request) (*API, error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Log(fmt.Sprintf("IP:%s connect %s with %s", r.RemoteAddr, r.RequestURI, r.Method))
		w.Header().Set("Content-Type", jsontype)
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(StatusInternalServerError.Status)
				json.NewEncoder(w).Encode(StatusInternalServerError.Body)
				utils.Err(fmt.Sprintf("IP:%s recover from "+string(debug.Stack()), r.RemoteAddr), err)
			}
			r.Body.Close()
		}()
		var j Request
		if err := json.NewDecoder(r.Body).Decode(&j); err != nil && err != io.EOF {
			w.WriteHeader(StatusBadRequest.Status)
			json.NewEncoder(w).Encode(StatusBadRequest.Body)
			utils.Err(fmt.Sprintf("IP:%s handle err", r.RemoteAddr), err)
			return
		}
		s, err := fn(&j, w, r)
		if err != nil {
			utils.Err(fmt.Sprintf("IP:%s", r.RemoteAddr), err)
		}
		if s != nil {
			w.WriteHeader(s.Status)
			if err := json.NewEncoder(w).Encode(s.Body); err != nil {
				w.WriteHeader(StatusInternalServerError.Status)
				json.NewEncoder(w).Encode(StatusInternalServerError.Body)
				utils.Err(fmt.Sprintf("IP:%s", r.RemoteAddr), err)
			}
		}
	}
}
