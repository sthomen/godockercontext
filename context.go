package main

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"github.com/fsnotify/fsnotify"
)

type Context struct {
	context string
	Channel chan string
}

func NewContext() *Context {
	context := new(Context)
	context.Channel = make(chan string)
	return context
}

func (self *Context) readContext(fn string) error {
	var conf struct {
		CurrentContext string
	}

	content, err := ioutil.ReadFile(fn)

	if err != nil{
		return err
	}

	err = json.Unmarshal(content, &conf)

	if err != nil {
		return err
	}

	if conf.CurrentContext != "" {
		self.context = conf.CurrentContext
	} else {
		self.context = "default"
	}

	return nil
}

func (self *Context) watch(fn string) {
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

	err = self.readContext(fn)

	if err != nil {
		log.Fatal(err)
		return
	}

	self.Channel <- self.context

	defer watcher.Close()

	for {
		select {
			case event, ok := <-watcher.Events:
				if !ok {
					log.Fatal(ok)
					return
				}

				if event.Op & fsnotify.Write == fsnotify.Write {
					err = self.readContext(fn)

					if err != nil {
						log.Fatal(err)
						return
					}

					self.Channel <- self.context
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
