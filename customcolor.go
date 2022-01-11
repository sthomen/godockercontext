package main

import (
	"strings"
	"errors"
	"fmt"

	"image/color"
)

// custom color holder for parsing custom color settings with flag
type customColors []struct {
	name string
	color color.RGBA
}

// not sure when this is used, but it's necessary for flag.Var
func (self *customColors) String() string {
	return "A custom color"
}

// Set method used by flag.Var
func (self *customColors) Set(value string) error {
	var parts = strings.Split(value, "=")

	if len(parts) != 2 {
		return errors.New("Color arguments need to be of the form name=color")
	}

	var c = new(color.RGBA)
	c.A = 255

	_, err := fmt.Sscanf(parts[1], "#%02x%02x%02x", &c.R, &c.G, &c.B)

	if err != nil {
		return err
	}

	*self = append(*self, struct {name string; color color.RGBA}{parts[0], *c})
	return nil
}
