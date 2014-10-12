package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	CONFIG   *Setting
	ROOT_WEB = os.Getenv("GOPATH") + "/src/github.com/fhacktory/ATYmgur/www/"
)

func main() {
	err := godotenv.Load()
	var imgur imgur
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	load_conf()

	imgur.CLIENT_ID = os.Getenv("CLIENT_ID")
	imgur.CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
	imgur.auth_url = "https://api.imgur.com/oauth2/authorize"
	imgur.token_url = "https://api.imgur.com/oauth2/token"
	imgur.init()
	if len(CONFIG.Album_key) == 0 {
		log.Println("No default album, creating one")
		imgur.create_album("ATYmgur", "Default folder of ATYmgur app", "hidden", "grid")
	}

	go folderWatcher([]string{"/tmp/foo", "/tmp/foo2"}, &imgur)
	startWeb()
}
