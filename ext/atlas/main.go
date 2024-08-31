package main

import (
	"image/png"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/atlas"
	"golang.org/x/image/colornames"
)

var (
	Atlas  atlas.Atlas
	group  = Atlas.MakeGroup()
	hiking = group.AddFile("hiking.png", png.Decode)
	gopher = Atlas.AddFile("thegopherproject.png", png.Decode)
)

func run() {
	cfg := opengl.WindowConfig{
		Title:  "Atlas example",
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := opengl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	Atlas.Pack()

	for !win.Closed() {
		if win.JustPressed(pixel.KeyEscape) {
			win.SetClosed(true)
		}

		win.Clear(colornames.Skyblue)

		hiking.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		gopher.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

		win.Update()
	}
}

func main() {
	opengl.Run(run)
}
