package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func stdWindow() *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  "PixTerra",
		Bounds: pixel.R(0, 0, 1024, 768),
		// VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	return win
}
