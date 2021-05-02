package main

import (
	"gobds/src/wss"
	"net/http"

	"github.com/gorilla/mux"
)

func router() {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	rapi := r.PathPrefix("/api").Subrouter()
	rwss := r.PathPrefix("/wss").Subrouter()
	rapi.HandleFunc("/server/{ServerID}")
	rwss.HandleFunc("/server/{ServerID}/{Session}", wss.ServerTerminal)

	http.ListenAndServe(":6623", r)
}
