package wss

import (
	"gobds/src/usefull"
	"net/http"
	"os"
)

var (
	err error
)

// Run ...
// start wss server
func Run() {
	// http.HandleFunc("/login/", echoLogin)
	http.HandleFunc("/ws/api/", echoAPI)
	http.HandleFunc("/ws/cmd/", echoCmd)
	http.HandleFunc("/ws/plg/", echoPlugin)
	// if err := http.ListenAndServeTLS(":6623", "gobds.cert", "gobds.key", nil); err != nil {
	// 	usefull.Err("wss server fail", err)
	// }
	usefull.Log("start....")
	if err = http.ListenAndServe(":6623", nil); err != nil {
		usefull.Err("ws server fail", err)
		os.Exit(10)
	}
}
