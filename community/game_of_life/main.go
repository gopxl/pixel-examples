package main

import (
	"flag"
	"math/rand"
	"time"

	"golang.org/x/image/colornames"

	"github.com/gopxl/pixel-examples/community/game_of_life/life"
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/imdraw"
)

var (
	size       *int
	windowSize *float64
	frameRate  *time.Duration
)

func init() {
	rand.Seed(time.Now().UnixNano())
	size = flag.Int("size", 5, "The size of each cell")
	windowSize = flag.Float64("windowSize", 800, "The pixel size of one side of the grid")
	frameRate = flag.Duration("frameRate", 33*time.Millisecond, "The framerate in milliseconds")
	flag.Parse()
}

func main() {
	opengl.Run(run)
}

func run() {

	cfg := opengl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, *windowSize, *windowSize),
		VSync:  true,
	}
	win, err := opengl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.Clear(colornames.White)

	// since the game board is square, rows and cols will be the same
	rows := int(*windowSize) / *size

	gridDraw := imdraw.New(nil)
	game := life.NewLife(rows, *size)
	tick := time.Tick(*frameRate)
	for !win.Closed() {
		// game loop
		select {
		case <-tick:
			gridDraw.Clear()
			game.A.Draw(gridDraw)
			gridDraw.Draw(win)
			game.Step()
		}
		win.Update()
	}
}
