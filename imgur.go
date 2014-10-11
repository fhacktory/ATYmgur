package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

func (i *imgur) upload_image(path string, title string) {
	reader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()
	imgdata, _ := ioutil.ReadAll(reader)
	img_base64 := base64.StdEncoding.EncodeToString(imgdata)

	resp, err := i.cl.PostForm("https://api.imgur.com/3/image",
		url.Values{
			"image":       {img_base64},
			"album":       {},
			"type":        {"base64"},
			"title":       {title},
			"description": {"Uploaded with ATYmgur"},
			"layout":      {"blog"}})
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}

func (i *imgur) create_album(name string, descr string, privacy string, layout string) {
	resp, err := i.cl.PostForm("https://api.imgur.com/3/album",
		url.Values{
			"title":       {name},
			"description": {descr},
			"privacy":     {privacy},
			"layout":      {layout}})
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}
