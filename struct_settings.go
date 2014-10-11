package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func encod(w http.ResponseWriter, r *http.Request) {
	treasure := make(map[string]string)
	treasure["path"] = "/usr/random"
	fmt.Println(treasure["path"])

	jsmap, _ := json.Marshal(treasure)
	fmt.Fprintf(w, string(jsmap), r.URL.Path[1:])
}
