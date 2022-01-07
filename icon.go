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
		color.RGBA{ 153, 184, 152, 255 }, // #99b898
		color.RGBA{ 254, 206, 168, 255 }, // #fecea8
		color.RGBA{ 255, 132, 124, 255 }, // #ff847c
		color.RGBA{ 232,  74,  95, 255 }, // #e84a5f
		color.RGBA{  42,  54,  59, 255 }, // #2a363b
	}

	var sum = 0

	for r := range context {
		sum += int(r)
	}

	return palette[sum % len(palette)]
}
