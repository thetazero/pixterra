package main

import (
	"math"

	"github.com/faiface/pixel"
)

func drawChunk(chunk [][]int, sheet []*pixel.Sprite, batch *pixel.Batch, chunkX int, chunkY int) {
	for y := 0; y < len(chunk); y++ {
		for x := 0; x < len(chunk[0]); x++ {
			position := chunkOffset(x, y)
			position.X += float64(chunkX) * 128.0
			position.Y += float64(chunkY) * 128.0
			r := pRandom(x, y)
			angle := 0.0
			switch {
			case r < 0.25:
				angle = math.Pi * 0.5
			case r < 0.5:
				angle = math.Pi
			case r < 0.75:
				angle = math.Pi * 1.5
			}
			it := sheet[chunk[y][x]]
			it.Draw(batch, pixel.IM.Scaled(pixel.ZV, 1.01).Rotated(pixel.ZV, angle).Moved(position))
		}
	}
}

func chunkOffset(x, y int) pixel.Vec {
	x1 := (float64(x) - 3.5) * 16.0
	y1 := (float64(y) - 3.5) * 16.0
	return pixel.V(x1, y1)
}
