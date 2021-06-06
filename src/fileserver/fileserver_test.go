package fileserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestFileServer(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {

	})
	s := httptest.NewServer(r)
	defer s.Close()

}
