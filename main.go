package main

//go:generate go run github.com/go-bindata/go-bindata/go-bindata -pkg $GOPACKAGE -o assets.go assets/

import (
	"os"
	"flag"
	"path/filepath"

	"github.com/getlantern/systray"
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

	state := NewState(*fn)

	systray.Run(state.onSystrayReady, nil)
}

type state struct {
	Filename string
	context *Context
}

func NewState(fn string) *state {
	state := new(state)
	state.Filename = fn
	return state
}

func (self *state) onSystrayReady() {

	icon, err := Asset("assets/icon.ico")

	if err == nil {
		systray.SetIcon(icon)
	}

	go menu()
	self.context = NewContext()
	go self.context.watch(self.Filename)

	for {
		var context = <-self.context.Channel
		systray.SetTitle(context)
		systray.SetTooltip(context)
	}
}

func menu() {
	quit := systray.AddMenuItem("Quit", "Stop monitoring context")

	for {
		select {
			case <-quit.ClickedCh:
				systray.Quit()
				return
		}
	}
}
