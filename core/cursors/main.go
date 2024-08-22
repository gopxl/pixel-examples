package main

import (
	"image/png"
	"os"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"golang.org/x/image/colornames"
)

func main() {
	opengl.Run(run)
}

func run() {
	cfg := opengl.WindowConfig{
		Title:  "Cursors",
		Bounds: pixel.R(0, 0, 1024, 768),
	}

	win, err := opengl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	f, err := os.Open("gopher-blushing.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	i, err := png.Decode(f)
	if err != nil {
		panic(err)
	}

	cursor := opengl.CreateCursorImage(i, pixel.ZV)
	win.SetCursor(cursor)

	for !win.Closed() {
		if win.JustReleased(pixel.KeyEscape) {
			win.SetClosed(true)
		}

		win.Clear(colornames.Skyblue)

		win.Update()
	}
}
