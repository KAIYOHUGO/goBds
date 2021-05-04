package main

import (
	"gobds/src/api"
	"gobds/src/wss"
	"net/http"

	"github.com/gorilla/mux"
)

func router() {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	{
		// api
		rapi := r.PathPrefix("/api").Subrouter()

		// session
		rapi.HandleFunc("/session", api.POSTSession).Methods("POST")
		rapi.HandleFunc("/session", api.DELETESession).Methods("DELETE")
		{
			// user
			ruser := rapi.PathPrefix("/user/{UserID}").Subrouter()
			ruser.HandleFunc("", api.POSTUser).Methods("POST")
			ruser.HandleFunc("", api.DELETEUser).Methods("DELETE")
			ruser.HandleFunc("/server", api.GETUserConfig).Methods("GET")
			ruser.HandleFunc("/server", api.PUTUserConfig).Methods("PUT")
			ruser.HandleFunc("/config", api.GETUserServer).Methods("GET")

		}
		rapi.HandleFunc("/servers/{ServerID}", api.GETServerFile).Methods("GET")
		rapi.HandleFunc("/server/{ServerID}", api.PUTServerFile).Methods("PUT")
	}
	// wss
	r.HandleFunc("wss/server/{ServerID}", wss.ServerTerminal)
	http.ListenAndServe(":6623", r)
}
