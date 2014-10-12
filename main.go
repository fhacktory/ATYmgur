package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
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

	//imgur.create_album("name", "descr", "hidden", "blog")
	//imgur.upload_image("/home/vayan/Downloads/foo.jpg", "foomage")

	folderWatcher([]string{"/tmp/foo", "/tmp/foo2"}, &imgur)
	startWeb()
}
