package main

import (
	"os"
	"flag"
	"path/filepath"
	"log"

//	"github.com/lxn/walk"
//	decl "github.com/lxn/walk/declarative"
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

func main() {
	var fn = flag.String("path", defaultPath(), "Path to the docker context config file")

	flag.Parse()

	context := make(chan string)

	go watch(context, *fn)

	for {
		log.Println("Context changed:", <-context)
	}
}
