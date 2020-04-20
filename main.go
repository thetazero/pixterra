package main

import (
	"fmt"
	"math"
	"time"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"

	"github.com/faiface/pixel/pixelgl"
)

type chunk [][]int

func run() {
	win := stdWindow()
	spriteSrc, spriteSheet := loadSpriteSheet("src/sprite.png", 16.0)

	var (
		camPos       = pixel.ZV
		camZoom      = 2.0
		camZoomSpeed = 1.15
		speed        = 100.0
	)
	var (
		frames = 0
		second = time.Tick(time.Second / 3)
	)

	var (
		last = time.Now()
	)

	world := [][]chunk{}
	for y := 0; y < 5; y++ {
		world = append(world, []chunk{})
		for x := 0; x < 5; x++ {
			world[y] = append(world[y], generateChunk(x-2, y-2))
		}
	}
	batch := pixel.NewBatch(&pixel.TrianglesData{}, spriteSrc)

	for y := 0; y < len(world); y++ {
		for x := 0; x < len(world[y]); x++ {
			drawChunk(world[y][x], spriteSheet, batch, x-2, y-2)
		}
	}
	// fmt.Println(world)
	chunk := pixel.V(0, 0)
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()
		cam := pixel.IM.Scaled(camPos, camZoom).Moved(win.Bounds().Center().Sub(camPos))
		win.SetMatrix(cam)
		if win.Pressed(pixelgl.KeyLeft) {
			camPos.X -= speed * dt
		}
		if win.Pressed(pixelgl.KeyRight) {
			camPos.X += speed * dt
		}
		if win.Pressed(pixelgl.KeyDown) {
			camPos.Y -= speed * dt
		}
		if win.Pressed(pixelgl.KeyUp) {
			camPos.Y += speed * dt
		}

		newChunk := calcChunk(camPos)
		if !chunk.Eq(newChunk) {
			chunk = newChunk
			updateWorld(world, spriteSheet, batch, chunk)
		}

		camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)
		win.Clear(colornames.Black)
		batch.Draw(win)
		spriteSheet[2].Draw(win, pixel.IM.Moved(camPos))
		win.Update()
		frames++
		select {
		case <-second:
			win.SetTitle(fmt.Sprintf("%s | FPS: %d", "PixTerra", frames*3))
			frames = 0
		default:
		}
	}
}

func calcChunk(p pixel.Vec) pixel.Vec {
	return pixel.V(math.Round(p.X/128), math.Round(p.Y/128))
}

func updateWorld(world [][]chunk, spriteSheet []*pixel.Sprite, batch *pixel.Batch, new pixel.Vec) {
	batch.Clear()
	for y := 0; y < len(world); y++ {
		for x := 0; x < len(world[y]); x++ {
			world[y][x] = generateChunk(int(new.X)+x-2, int(new.Y)+y-2)
			drawChunk(world[y][x], spriteSheet, batch, x-2+int(new.X), y-2+int(new.Y))
		}
	}
}

func main() {
	pixelgl.Run(run)
}
