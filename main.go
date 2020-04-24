package main

import (
	"fmt"
	"math"
	"time"

	"golang.org/x/image/colornames"

	"github.com/faiface/pixel"
	"github.com/ojrac/opensimplex-go"

	"github.com/faiface/pixel/pixelgl"
)

type chunk [][]int

var spriteSrc pixel.Picture

func run() {
	win := stdWindow()
	src, spriteSheet := loadSpriteSheet("src/sprite.png", 16.0)
	spriteSrc = src
	var (
		camPos       = pixel.ZV
		camZoom      = 2.0
		camZoomSpeed = 1.15
		speed        = 300.0
	)
	var (
		frames = 0
		second = time.Tick(time.Second / 3)
	)

	var (
		last = time.Now()
	)

	world := [][]chunk{}
	w := 7
	batches := [][]*pixel.Batch{}

	for y := 0; y < w; y++ {
		world = append(world, []chunk{})
		batches = append(batches, []*pixel.Batch{})
		for x := 0; x < w; x++ {
			world[y] = append(world[y], generateChunk(x-w/2, y-w/2))
			batches[y] = append(batches[y], pixel.NewBatch(&pixel.TrianglesData{}, spriteSrc))
		}
	}

	for y := 0; y < len(world); y++ {
		for x := 0; x < len(world[y]); x++ {
			drawChunk(world[y][x], spriteSheet, batches[y][x], x-w/2, y-w/2)
		}
	}
	// fmt.Println(world)
	chunk := pixel.V(0, 0)
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()
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
		camZoom *= math.Pow(camZoomSpeed, win.MouseScroll().Y)
		cam := pixel.IM.Scaled(camPos, camZoom).Moved(win.Bounds().Center().Sub(camPos))
		win.SetMatrix(cam)

		newChunk := calcChunk(camPos)
		if !chunk.Eq(newChunk) {
			updateWorld(world, spriteSheet, batches, chunk, newChunk)
			chunk = newChunk
		}

		win.Clear(colornames.Black)
		drawBatches(batches, win)
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

func updateWorld(world [][]chunk, spriteSheet []*pixel.Sprite, batches [][]*pixel.Batch, old pixel.Vec, new pixel.Vec) {
	w := len(world)
	dir := old.Sub(new)
	fmt.Println(dir)
	if dir.Y == 1 {
		for y := w - 1; y > 0; y-- {
			for x := 0; x < len(world[0]); x++ {
				world[y][x] = world[y-1][x]
				batches[y][x] = batches[y-1][x]
			}
		}
		for x := 0; x < len(world[0]); x++ {
			world[0][x] = nil
			batches[0][x] = nil
		}
	} else if dir.Y == -1 {
		for y := 0; y < w-1; y++ {
			for x := 0; x < len(world[0]); x++ {
				world[y][x] = world[y+1][x]
				batches[y][x] = batches[y+1][x]
			}
		}
		for x := 0; x < len(world[0]); x++ {
			world[w-1][x] = nil
			batches[w-1][x] = nil
		}
	}

	if dir.X == 1 {
		for x := len(world[0]) - 1; x > 0; x-- {
			for y := 0; y < w; y++ {
				world[y][x] = world[y][x-1]
				batches[y][x] = batches[y][x-1]
			}
		}
		for y := 0; y < w; y++ {
			world[y][0] = nil
			batches[y][0] = nil
		}
	} else if dir.X == -1 {
		for x := 0; x < len(world[0])-1; x++ {
			for y := 0; y < w; y++ {
				world[y][x] = world[y][x+1]
				batches[y][x] = batches[y][x+1]
			}
		}
		for y := 0; y < w; y++ {
			world[y][len(world[0])-1] = nil
			batches[y][len(world[0])-1] = nil
		}
	}

	for y := 0; y < len(world); y++ {
		for x := 0; x < len(world[y]); x++ {
			if world[y][x] == nil {
				world[y][x] = generateChunk(int(new.X)+x-w/2, int(new.Y)+y-w/2)
				batches[y][x] = pixel.NewBatch(&pixel.TrianglesData{}, spriteSrc)
				drawChunk(world[y][x], spriteSheet, batches[y][x], x-w/2+int(new.X), y-w/2+int(new.Y))
			}
		}
	}
}

func drawBatches(batches [][]*pixel.Batch, win *pixelgl.Window) {
	for y := 0; y < len(batches); y++ {
		for x := 0; x < len(batches[0]); x++ {
			batches[y][x].Draw(win)
		}
	}
}

var simplex opensimplex.Noise

func main() {
	simplex = opensimplex.New(69)
	pixelgl.Run(run)
}
