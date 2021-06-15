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

func init() {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	{
		// api
		rapi := r.PathPrefix("/api").Subrouter()

		// session
		rapi.HandleFunc("/session", api.Wrapper(api.POSTSession)).Methods(http.MethodPost)
		rapi.HandleFunc("/session", api.Wrapper(api.DELETESession)).Methods(http.MethodDelete)
		rapi.HandleFunc("/user", api.Wrapper(api.POSTUser)).Methods(http.MethodPost)
		rapi.HandleFunc("/user", api.Wrapper(api.DELETEUser)).Methods(http.MethodDelete)
		rapi.HandleFunc("/server", api.Wrapper(api.PostServer)).Methods(http.MethodPost)
		rapi.HandleFunc("/server", api.Wrapper(api.DeleteServer)).Methods(http.MethodDelete)
		{
			// user
			ruser := rapi.PathPrefix("/user/{UserID}").Subrouter()
			ruser.HandleFunc("/config", api.GETUserConfig).Methods("GET")
			ruser.HandleFunc("/config", api.Wrapper(api.PUTUserConfig)).Methods("PUT")
			ruser.HandleFunc("/servers", api.Wrapper(api.GETUserServers)).Methods("GET")

		}
		{
			rserver := rapi.PathPrefix("/server/{ServerID}").Subrouter()
			rserver.HandleFunc("/input", api.Wrapper(api.POSTServerInput)).Methods(http.MethodPost)
			{
				// server
				rserverfile := rserver.PathPrefix("/file").Subrouter()
				rserverfile.HandleFunc("", api.Wrapper(api.GETServerFile)).Methods(http.MethodGet)
				rserverfile.HandleFunc("/{Path}", api.Wrapper(api.GETServerFile)).Methods(http.MethodGet)
				rserverfile.HandleFunc("/{Path}", api.Wrapper(api.PUTServerFile)).Methods(http.MethodPut)
				rserverfile.HandleFunc("/{Path}:{Type}", api.Wrapper(api.POSTServerFile)).Methods(http.MethodPost)
				rserverfile.HandleFunc("/{Path}", api.Wrapper(api.DELETEServerFile)).Methods(http.MethodDelete)
				rserverfile.HandleFunc("/{Path}:{NewPath}", api.Wrapper(api.PATCHServerFile)).Methods(http.MethodPatch)
			}
		}
	}
	// wss
	r.HandleFunc("/wss/server/{ServerID}/terminal/{SessionID}", wss.ServerTerminal)
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	utils.Log(fmt.Sprintf("Run on localhost:%d", l.Addr().(*net.TCPAddr).Port))
	panic(http.Serve(l, r))
}
