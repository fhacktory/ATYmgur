package main

import (
	"fmt"
	"gopkg.in/fsnotify.v1"
	"io/ioutil"
)

func initFolder(folderPath string) /*(fileInfo os.FileInfo)*/ {

	//func ReadDir(dirname string) ([]os.FileInfo, error)

	dir, _ := ioutil.ReadDir(folderPath)
	for _, f := range dir {
		fmt.Println(f.Name())
	}
	/*chann := make(chan string)
	go func() {
		filepath.Walk(folderPath, func(path string, fileInfo os.FileInfo, _ error) (err error) {
			chann <- path
			return
		})
		defer close(chann)
	}()
	fmt.Println(fileInfo.Name)

	return fileInfo.Name()*/
}

func folderWatcher() {
	var i int
	watcher, err := fsnotify.NewWatcher()
	foldersNamesArray := []string{"/tmp/foo", "/tmp/foo2"}

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
