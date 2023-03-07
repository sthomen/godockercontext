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

	switch conf.CurrentContext {
		case "":
			self.context = "default"
		default:
			self.context = conf.CurrentContext
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
					log.Fatal("watcher.Events ok? ", ok)
				}

				if event.Has(fsnotify.Remove) || event.Has(fsnotify.Rename) {
					watcher.Remove(fn)
					watcher.Add(fn)
				}

				// Immediately read the context even after a Remove or Rename
				// event has happened because fsnotify is (presumably) too slow
				// to react to the quick remove-create-write of the docker cli
				// when the context is being changed.

				err = self.readContext(fn)

				if err != nil {
					log.Println("self.readContext(fn) err = ", err)
				}

				self.Channel <- self.context

			case err, ok := <-watcher.Errors:
				if !ok {
					log.Fatal("watcher.Events ok? ", ok)
				}

				log.Fatal("watcher.Events err = ", err)
		}
	}
}
