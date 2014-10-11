package main

import (
	"io/ioutil"
	"log"
	"net/url"
)

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
