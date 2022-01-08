package main

import (
	"image"
	"image/draw"
	"image/color"
	ico "git.shangtai.net/staffan/go-ico"
)

func generateIconFromContextString(context string) []byte {
	icon := new(ico.Icon)

	img := image.NewRGBA(image.Rect(0,0,64,64))
	draw.Draw(img, img.Bounds(), &image.Uniform{colorFromContextString(context)}, image.ZP, draw.Src)

	icon.AddImage(img)

	ico, err := icon.Encode()

	if err != nil {
		return nil
	}

	return ico
}

func colorFromContextString(context string) color.RGBA {
	palette := []color.RGBA{
		color.RGBA{230,  25,  75, 255},
		color.RGBA{ 60, 180,  75, 255},
		color.RGBA{255, 225,  25, 255},
		color.RGBA{ 67,  99, 216, 255},
		color.RGBA{245, 130,  49, 255},
		color.RGBA{145,  30, 180, 255},
		color.RGBA{ 70, 240, 240, 255},
		color.RGBA{240,  50, 230, 255},
		color.RGBA{188, 246,  12, 255},
		color.RGBA{250, 190, 190, 255},
		color.RGBA{  0, 128, 128, 255},
		color.RGBA{230, 190, 255, 255},
		color.RGBA{154,  99,  36, 255},
		color.RGBA{255, 250, 200, 255},
		color.RGBA{128,   0,   0, 255},
		color.RGBA{170, 255, 195, 255},
		color.RGBA{128, 128,   0, 255},
		color.RGBA{255, 216, 177, 255},
		color.RGBA{  0,   0, 117, 255},
		color.RGBA{128, 128, 128, 255},
		color.RGBA{255, 255, 255, 255},
		color.RGBA{  0,   0,   0, 255},
	}

	var sum = 0

	for r := range context {
		sum += int(r)
	}

	return palette[sum % len(palette)]
}
