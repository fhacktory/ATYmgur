package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func encod(w http.ResponseWriter) {
	treasure := make(map[string]string)
	treasure["path"] = "/usr/random"

	jsmap, _ := json.Marshal(treasure)
	fmt.Fprintf(w, string(jsmap))
}
