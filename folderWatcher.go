package main

import (
	"fmt"
	"gopkg.in/fsnotify.v1"
)

func folderWatcher() {
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
				fmt.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println("modified file:", event.Name)
				}
			case err := <-watcher.Errors:
				fmt.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("/tmp/foo")
	err = watcher.Add("/tmp/foo2")

	if err != nil {
		fmt.Println(err)
	}
	<-done
}
