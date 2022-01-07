package main

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
	// an icon must exist before a menu item is added, so lets just use the default
	systray.SetIcon(generateIconFromContextString("default"))

	// display the current context as a disabled entry in the menu,
	// default to "default" here, it will be immediately overwritten by the
	// loop below, but there has to be something.
	current := systray.AddMenuItem("default", "Current context")
	current.Disable()

	// launch a goroutine to handle clicks
	go menuClickHandler()

	// file watcher goroutine, feeds the current context into its channel
	self.context = NewContext()
	go self.context.watch(self.Filename)

	// wait for a context from the channel and set it in the tray
	for {
		var context = <-self.context.Channel

		systray.SetIcon(generateIconFromContextString(context))
		current.SetTitle(context)
		systray.SetTitle(context)
		systray.SetTooltip(context)
	}
}

func menuClickHandler() {
	quit := systray.AddMenuItem("Quit", "Stop monitoring context")

	for {
		select {
			case <-quit.ClickedCh:
				systray.Quit()
				return
		}
	}
}
