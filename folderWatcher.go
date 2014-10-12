package main

import (
	"fmt"
	"gopkg.in/fsnotify.v1"
	"io/ioutil"
	"log"
)

func initFolder(folderPath string, img *imgur) {
	var isFile bool
	dir, _ := ioutil.ReadDir(folderPath)

	log.Println("Uploading content of folder ", folderPath)

	for _, f := range dir {
		isFile = fileCheck(f.Name())
		if isFile == true {
			img.upload_image(folderPath+"/"+f.Name(), f.Name())

		} else {
			log.Println(f.Name() + " Extension not valid, upload an image pls")
		}
	}
}

func folderWatcher(foldersNamesArray []string, img *imgur) {
	var i int
	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		fmt.Println(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op == fsnotify.Create {
					log.Println("New image detected")
					img.upload_image(event.Name, "atymgur")
				}
			case err := <-watcher.Errors:
				fmt.Println("error:", err)
			}
		}
	}()

	for i = 0; i < len(foldersNamesArray); i++ {
		go initFolder(foldersNamesArray[i], img)
		err = watcher.Add(foldersNamesArray[i])
	}

	if err != nil {
		fmt.Println(err)
	}
	<-done
}
