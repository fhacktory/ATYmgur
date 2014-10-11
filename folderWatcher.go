package main

import (
	"fmt"
	"gopkg.in/fsnotify.v1"
)

func folderWatcher() {
	var i int
	watcher, err := fsnotify.NewWatcher()
	folderNamesArray := []string{"/tmp/foo", "/tmp/foo2"}

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

	for i = 0; i < len(folderNamesArray); i++ {
		err = watcher.Add(folderNamesArray[i])
	}

	if err != nil {
		fmt.Println(err)
	}
	<-done
}
