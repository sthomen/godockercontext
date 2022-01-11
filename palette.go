package main

import "image/color"

var defaultColors = []color.RGBA{
	color.RGBA{255, 255, 255, 255},
	color.RGBA{  0,   0, 117, 255},
	color.RGBA{  0, 128, 128, 255},
	color.RGBA{ 60, 180,  75, 255},
	color.RGBA{ 67,  99, 216, 255},
	color.RGBA{ 70, 240, 240, 255},
	color.RGBA{128,   0,   0, 255},
	color.RGBA{128, 128,   0, 255},
	color.RGBA{128, 128, 128, 255},
	color.RGBA{145,  30, 180, 255},
	color.RGBA{154,  99,  36, 255},
	color.RGBA{170, 255, 195, 255},
	color.RGBA{188, 246,  12, 255},
	color.RGBA{230,  25,  75, 255},
	color.RGBA{230, 190, 255, 255},
	color.RGBA{240,  50, 230, 255},
	color.RGBA{245, 130,  49, 255},
	color.RGBA{250, 190, 190, 255},
	color.RGBA{255, 216, 177, 255},
	color.RGBA{255, 225,  25, 255},
	color.RGBA{255, 250, 200, 255},
}

type Palette struct {
	colors []struct {
		name string
		color color.RGBA
	}
}

func (self *Palette) Set(name string, hue color.RGBA) {
	var entry = struct {name string; color color.RGBA}{name: name, color: hue}

	for index, color := range self.colors {
		if (color.name == name) {
			self.colors[index] = entry
			return
		}
	}

	self.colors = append(self.colors, entry)
}

func (self *Palette) GetColor(name string) color.RGBA {
	for _, entry := range self.colors {
		if entry.name == name {
			return entry.color
		}
	}

	return defaultColors[nameToIndex(name) % len(defaultColors)]
}

func nameToIndex(name string) int {
	var sum = 0

	for r := range name {
		sum += int(r)
	}

	return sum
}
