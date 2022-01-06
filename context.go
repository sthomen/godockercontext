package main

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"github.com/fsnotify/fsnotify"
)

func readContext(fn string) (string, error) {
	var conf struct {
		CurrentContext string
	}

	content, err := ioutil.ReadFile(fn)

	if err != nil{
		return "", err
	}

	err = json.Unmarshal(content, &conf)

	if err != nil {
		return "", err
	}

	if conf.CurrentContext == "" {
		return "default", nil
	}

	return conf.CurrentContext, nil
}

func watch(output chan string, fn string) {
	log.Println(fn)

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
		return
	}

	err = watcher.Add(fn)

	if err != nil {
		log.Fatal(err)
		return
	}

	context, err := readContext(fn)

	if err != nil {
		log.Fatal(err)
		return
	}

	output <- context

	defer watcher.Close()

	for {
		select {
			case event, ok := <-watcher.Events:
				if !ok {
					log.Fatal(ok)
					return
				}

				if event.Op & fsnotify.Write == fsnotify.Write {
					context, err := readContext(fn)

					if err != nil {
						log.Fatal(err)
						return
					}

					output <- context
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					log.Fatal(ok)
					return
				}

				log.Fatal(err)
		}
	}
}
