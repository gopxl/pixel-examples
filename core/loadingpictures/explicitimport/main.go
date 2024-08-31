package main

import (
	"image/jpeg"
	"image/png"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"golang.org/x/image/colornames"
)

func main() {
	opengl.Run(run)
}

func run() {
	cfg := opengl.WindowConfig{
		Title:  "Image Loading",
		Bounds: pixel.R(0, 0, 1024, 768),
	}

	win, err := opengl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	p1, err := pixel.PictureDataFromFile("../hiking.png", png.Decode)
	if err != nil {
		panic(err)
	}

	p2, err := pixel.PictureDataFromFile("../hiking.jpeg", jpeg.Decode)
	if err != nil {
		panic(err)
	}

	spr1 := pixel.NewSprite(p1, p1.Bounds())
	spr2 := pixel.NewSprite(p2, p2.Bounds())

	for !win.Closed() {
		if win.JustReleased(pixel.KeyEscape) {
			win.SetClosed(true)
		}

		win.Clear(colornames.Skyblue)

		spr1.Draw(win, pixel.IM.Moved(win.Bounds().Center().Add(pixel.V(0, 100))))
		spr2.Draw(win, pixel.IM.Moved(win.Bounds().Center().Add(pixel.V(0, -100))))

		win.Update()
	}
}
