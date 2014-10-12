package main

import (
	"strings"
)

func fileCheck(fileName string) bool {
	var i int
	var matched bool
	imgExt := []string{
		".jpeg",
		".jpg",
		".JPG",
		".JPEG",
		".png",
		".PNG",
		".gif",
		".GIF",
	}

	for i = 0; i < len(imgExt); i++ {
		if strings.Contains(fileName, imgExt[i]) {
			matched = true
			i = len(imgExt)
		} else {
			matched = false
		}
	}

	return matched
}
