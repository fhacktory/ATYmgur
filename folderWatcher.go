package main

import (
	"fmt"
	"gopkg.in/fsnotify.v1"
	"io/ioutil"
	"os"
)

func initFolder(folderPath string) (fileInfo os.FileInfo) {
	dir, _ := ioutil.ReadDir(folderPath)
	for _, f := range dir {
		fmt.Println(f.Name())
	}

	return fileInfo
}

func folderWatcher() {
	var i int
	watcher, err := fsnotify.NewWatcher()
	foldersNamesArray := []string{"/tmp/foo", "/tmp/foo2", "/tmp/foo3"}

	if err != nil {
		fmt.Println(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				fmt.Println("error:", err)
			}
		}
	}()

	for i = 0; i < len(foldersNamesArray); i++ {
		fileInfo := initFolder(foldersNamesArray[i])
		// Yann's plug there
		fileInfo = fileInfo
		err = watcher.Add(foldersNamesArray[i])
	}

	if err != nil {
		fmt.Println(err)
	}
	<-done
}
