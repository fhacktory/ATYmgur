package main

//Web Server in Progress : #BetterthanApache
//m
import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, Welcome to %s!", r.URL.Path[1:])
}

func startWeb() {
	fmt.Println("Starting WebServer")
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/settings", SettingsHandler)
	r.HandleFunc("/callback", CallbackHandler)
	http.ListenAndServe(":8181", nil)
}
