package main

import (
	"os"
	"path/filepath"
	"flag"
	"log"
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"encoding/json"

	"github.com/lxn/walk"
	decl "github.com/lxn/walk/declarative"
)

func defaultPath() string {
	homedir, err := os.UserHomeDir()

	// default to SOMETHING
	if err != nil {
		homedir = "/"
	}

	// FromSlash because filepath somehow doesn't handle windows directory
	// separators directly.
	path := filepath.FromSlash(filepath.Join(homedir, ".docker", "config.json"))

	return path
}


func readContext(fn string) string {
	var conf struct {
		CurrentContext string
	}

	content, err := ioutil.ReadFile(fn)

	if err != nil{
		log.Fatal(err)
		return ""
	}

	err = json.Unmarshal(content, &conf)

	if err != nil {
		log.Fatal(err)
		return ""
	}

	return conf.CurrentContext
}

func main() {
	var fn = flag.String("path", defaultPath(), "Path to the docker context config file")

	flag.Parse()

	log.Println("Watching", *fn)

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
		return
	}

	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}

					if event.Op & fsnotify.Write == fsnotify.Write {
						log.Println("Context:", readContext(event.Name))
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}

					log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(*fn)

	if err != nil {
		log.Fatal(err)
	}

	<-done
}
