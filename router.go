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
	r.HandleFunc("/wss/server/{ServerID}/terminal/{SessionID}", wss.ServerTerminal)
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}

	utils.Log(fmt.Sprintf("Run on localhost:%d", l.Addr().(*net.TCPAddr).Port))
	panic(http.Serve(l, r))
}
