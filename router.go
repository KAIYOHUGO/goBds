package main

import (
	"fmt"
	"gobds/src/api"
	"gobds/src/utils"
	"gobds/src/wss"
	"net"
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
		rapi.HandleFunc("/session", api.POSTSession).Methods(http.MethodPost)
		rapi.HandleFunc("/session", api.DELETESession).Methods(http.MethodDelete)
		rapi.HandleFunc("/user", api.POSTUser).Methods("POST")
		rapi.HandleFunc("/user", api.DELETEUser).Methods("DELETE")
		{
			// user
			ruser := rapi.PathPrefix("/user/{UserID}").Subrouter()
			ruser.HandleFunc("/server", api.GETUserConfig).Methods(http.MethodGet)
			ruser.HandleFunc("/server", api.PUTUserConfig).Methods(http.MethodPut)
			ruser.HandleFunc("/config", api.GETUserServer).Methods(http.MethodGet)

		}
		rapi.HandleFunc("/servers/{ServerID}", api.GETServerFile).Methods(http.MethodGet)
		rapi.HandleFunc("/server/{ServerID}", api.PUTServerFile).Methods(http.MethodPut)
	}
	// wss
	r.HandleFunc("/wss/server/{ServerID}/terminal/{SessionID}", wss.ServerTerminal)
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	utils.Log(fmt.Sprintf("Run on localhost:%d", l.Addr().(*net.TCPAddr).Port))
	panic(http.Serve(l, r))
}
