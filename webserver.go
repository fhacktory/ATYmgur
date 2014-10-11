package main

//Web Server in Progress : #BetterthanApache

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, Welcome to Home!", r.URL.Path[1:])
}

func Settings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, Set your %s!", r.URL.Path[1:])
}

func Callback(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, %s page!", r.URL.Path[1:])
}

func startWeb() {
	r := mux.NewRouter()
	fmt.Println("Starting WebServer")
	r.HandleFunc("/", Home)
	r.HandleFunc("/settings", Settings)
	r.HandleFunc("/callback", Callback)
	http.Handle("/", r)
	http.ListenAndServe(":8181", nil)
	fmt.Println("Stoping WebServer")
}
