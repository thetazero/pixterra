package main

import (
	"image/png"
	"os"

	"github.com/faiface/pixel"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func loadSprite(path string) *pixel.Sprite {
	pic, err := loadPicture(path)
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(pic, pic.Bounds())
	return sprite
}

func loadSpriteSheet(path string, size float64) (pixel.Picture, []*pixel.Sprite) {
	src, err := loadPicture(path)
	if err != nil {
		panic(err)
	}

	var sheet []*pixel.Sprite
	for x := src.Bounds().Min.X; x < src.Bounds().Max.X; x += size {
		for y := src.Bounds().Min.Y; y < src.Bounds().Max.Y; y += size {
			frame := pixel.R(x, y, x+size, y+size)
			sheet = append(sheet, pixel.NewSprite(src, frame))
		}
	}
	return src, sheet
}
