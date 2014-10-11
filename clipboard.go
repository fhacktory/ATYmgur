package main

import (
	"github.com/atotto/clipboard"
)

func add_img_to_clipboard(url string) {
	clipboard.WriteAll(url)
}
