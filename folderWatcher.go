package main

import (
	"fmt"
	"gopkg.in/fsnotify.v1"
	"io/ioutil"
)

// Checks if folder already got images in it
func initFolder(folderPath string, img *imgur) {
	var isFile bool
	dir, _ := ioutil.ReadDir(folderPath)
	for _, f := range dir {
		isFile = fileCheck(f.Name())
		if isFile == true {
			fmt.Println(f.Name())
			img.upload_image(folderPath+"/"+f.Name(), "foobarfoobar")
		} else {
			fmt.Println(f.Name() + " Extension not valid, upload an image pls")
		}
	}
}

// Loops on foldersNamesArray and add watcher to every one of them
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
					img.upload_image(event.Name, "foobarfoobar")
				}
			case err := <-watcher.Errors:
				fmt.Println("error:", err)
			}
		}
	}()

	for i = 0; i < len(foldersNamesArray); i++ {
		initFolder(foldersNamesArray[i], img)
		err = watcher.Add(foldersNamesArray[i])
	}

	if err != nil {
		fmt.Println(err)
	}
	<-done
}
