package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Setting struct {
	Album_key string
	Key       string
	filename  string
}

func save_conf() {
	stringjson, err := json.Marshal(CONFIG)
	if err != nil {
		log.Println("save conf: ", err)
	}
	file, err := os.Create(CONFIG.filename)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(" Write to file : " + CONFIG.filename)
	n, err := io.WriteString(file, string(stringjson))
	if err != nil {
		fmt.Println(n, err)
	}
	file.Close()
}

func load_conf() {
	temp := new(Setting)
	temp.filename = "config.json"

	file, err := ioutil.ReadFile(temp.filename)
	if err != nil {
		log.Println("open config: ", err)
	}

	if err = json.Unmarshal(file, temp); err != nil {
		log.Println("parse config: ", err)
	}
	CONFIG = temp
}

func encod(w http.ResponseWriter) {
	treasure := make(map[string]string)
	treasure["path"] = "/usr/random"

	jsmap, _ := json.Marshal(treasure)
	fmt.Fprintf(w, string(jsmap))
}
