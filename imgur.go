package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

type AnswerRequest struct {
	Data    ImageMetaData
	Success bool
}

type ImageMetaData struct {
	Id          string
	Title       string
	Description string
	Datetime    int64
	Type        string
	Animated    bool
	Width       int64
	Height      int64
	Size        int64
	Views       int64
	Link        string
}

func (i *imgur) upload_image(path string, title string) string {
	var imMeta AnswerRequest

	log.Println("Starting upload : ", path)
	reader, err := os.Open(path)
	if err != nil {
		log.Println("Error open file to upload :", err)
	}
	defer reader.Close()
	imgdata, _ := ioutil.ReadAll(reader)
	img_base64 := base64.StdEncoding.EncodeToString(imgdata)

	resp, err := i.cl.PostForm("https://api.imgur.com/3/image",
		url.Values{
			"image":       {img_base64},
			"album":       {CONFIG.Album_key},
			"type":        {"base64"},
			"title":       {title},
			"description": {"Uploaded with ATYmgur"},
			"layout":      {"blog"}})
	if err != nil {
		log.Println("Error send new image :", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &imMeta)
	if err != nil {
		log.Println("Error upload : ", err)
	}
	add_img_to_clipboard(imMeta.Data.Link)
	log.Println("Upload finished : ", path)
	if !imMeta.Success {
		log.Println("Upload wasn't successful : ", string(body))
	}
	return imMeta.Data.Link
}

func (i *imgur) create_album(name string, descr string, privacy string, layout string) {
	var albMeta AnswerRequest

	resp, err := i.cl.PostForm("https://api.imgur.com/3/album",
		url.Values{
			"title":       {name},
			"description": {descr},
			"privacy":     {privacy},
			"layout":      {layout}})
	if err != nil {
		log.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &albMeta)
	if err != nil {
		log.Println("Error upload : ", err)
	}
	CONFIG.Album_key = albMeta.Data.Id
	save_conf()
}
