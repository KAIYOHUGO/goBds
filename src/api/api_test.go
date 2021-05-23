package api

import (
	"bytes"
	"encoding/json"
	"gobds/src/config"
	"gobds/src/database"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetToken(t *testing.T) {
	s, err := GetToken("Bearer paula0623")
	if err != nil {
		t.Fatal(err)
	}
	if s != "paula0623" {
		t.Fatal("not the same")
	}
	t.Log(s)
}

func TestAPI(t *testing.T) {

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	{
		// api
		rapi := r.PathPrefix("/api").Subrouter()

		// session
		rapi.HandleFunc("/session", POSTSession).Methods("POST")
		rapi.HandleFunc("/session", DELETESession).Methods("DELETE")
		rapi.HandleFunc("/user", POSTUser).Methods("POST")
		rapi.HandleFunc("/user", DELETEUser).Methods("DELETE")
		{
			// user
			ruser := rapi.PathPrefix("/user/{UserID}").Subrouter()
			ruser.HandleFunc("/server", GETUserConfig).Methods("GET")
			ruser.HandleFunc("/server", PUTUserConfig).Methods("PUT")
			ruser.HandleFunc("/config", GETUserServer).Methods("GET")

		}
		rapi.HandleFunc("/servers/{ServerID}", GETServerFile).Methods("GET")
		rapi.HandleFunc("/server/{ServerID}", PUTServerFile).Methods("PUT")
	}
	s := httptest.NewServer(r)
	defer s.Close()
	database.Write(database.DB["account"], "paula", config.Account{
		Name: "paula",
		// 12345678
		Password: "7c222fb2927d828af22f592134e8932480637c0d",
	})

	var session string
	// funcs
	testsession := func(t *testing.T) {
		var session string
		{
			req, err := http.NewRequest(http.MethodPost, s.URL+"/api/session", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetBasicAuth("paula", "12345678")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			session, err = GetToken(resp.Header.Get("Authorization"))
			if err != nil {
				t.Fatal(err)
			}
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			t.Log(string(b))
			t.Log(resp.StatusCode)
			if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
				t.Fatal(resp.StatusCode)
			}
			t.Log(session)
		}
		req, err := http.NewRequest(http.MethodDelete, s.URL+"/api/session", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer "+session)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(resp.StatusCode)
		if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
			t.Fatal(resp.StatusCode)
		}
	}
	testuser := func(t *testing.T) {
		{
			req, err := http.NewRequest(http.MethodPost, s.URL+"/api/session", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetBasicAuth("paula", "12345678")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			session, err = GetToken(resp.Header.Get("Authorization"))
			if err != nil {
				t.Fatal(err)
			}
			t.Log(session)
		}
		{
			b, err := json.Marshal(&ReqUser{
				Name:     "Sorry",
				Password: "Paula0623",
			})
			if err != nil {
				t.Fatal(err)
			}
			req, err := http.NewRequest(http.MethodPost, s.URL+"/api/user", bytes.NewBuffer(b))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Set("Authorization", "Bearer "+session)
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			b, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}
			t.Log(string(b))
		}
		{
			req, err := http.NewRequest(http.MethodPost, s.URL+"/api/session", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetBasicAuth("Sorry", "Paula0623")
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()
			session, err = GetToken(resp.Header.Get("Authorization"))
			if err != nil {
				t.Fatal(err)
			}
			t.Log(session)
		}
	}
	t.Run("session", testsession)
	t.Run("user", testuser)
}
