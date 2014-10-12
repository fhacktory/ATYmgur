package main

//Web Server in Progress : #BetterthanApache

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(ROOT_WEB + "index.html")
	t.Execute(w, nil)
}

func Settings(w http.ResponseWriter, r *http.Request) {
	encod(w)
}

func Callback(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, callback page!")
}

func startWeb() {
	r := mux.NewRouter()
	fmt.Println("Starting WebServer")
	r.HandleFunc("/", Home)
	r.HandleFunc("/settings", Settings)
	r.HandleFunc("/callback", Callback)
	r.PathPrefix("/assets/").Handler(http.FileServer(http.Dir(ROOT_WEB)))
	http.Handle("/", r)
	http.ListenAndServe(":8181", nil)
	fmt.Println("Stoping WebServer")
}
