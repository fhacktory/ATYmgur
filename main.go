package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	ROOT_WEB = os.Getenv("GOPATH") + "/src/github.com/fhacktory/ATYmgur/www/"
)

func main() {
	err := godotenv.Load()
	var imgur imgur
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	imgur.CLIENT_ID = os.Getenv("CLIENT_ID")
	imgur.CLIENT_SECRET = os.Getenv("CLIENT_SECRET")
	imgur.auth_url = "https://api.imgur.com/oauth2/authorize"
	imgur.token_url = "https://api.imgur.com/oauth2/token"
	imgur.init()

	go folderWatcher([]string{"/home/vayan/up_to_imgur"}, &imgur)
	startWeb()
}
