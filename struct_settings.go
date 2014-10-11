package main

import (
	"encoding/json"
	"fmt"
)

func encod() {
	treasure := make(map[string]string)
	treasure["path"] = "/usr/random"
	fmt.Println(treasure["path"])

	jsmap, _ := json.Marshal(treasure)
	fmt.Println("jsmap : ", string(jsmap))
}
