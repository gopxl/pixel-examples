package main

import (
	"fmt"
	"log"
	"time"

	_ "image/png"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/atlas"
	"github.com/gopxl/pixelui/v2"
	"github.com/inkyblackness/imgui-go/v4"
	"golang.org/x/image/colornames"
)

var (
	Atlas           atlas.Atlas
	angry           = Atlas.AddFile("emoji/gopher-angry.png")
	atPeace         = Atlas.AddFile("emoji/gopher-at-peace.png")
	blushing        = Atlas.AddFile("emoji/gopher-blushing.png")
	coldSweat       = Atlas.AddFile("emoji/gopher-cold-sweat.png")
	confused        = Atlas.AddFile("emoji/gopher-confused.png")
	cryingRiver     = Atlas.AddFile("emoji/gopher-crying-river.png")
	crying          = Atlas.AddFile("emoji/gopher-crying.png")
	dead            = Atlas.AddFile("emoji/gopher-dead.png")
	expressionless  = Atlas.AddFile("emoji/gopher-expressionless.png")
	facepalm        = Atlas.AddFile("emoji/gopher-facepalm.png")
	happy           = Atlas.AddFile("emoji/gopher-happy.png")
	heartEyes       = Atlas.AddFile("emoji/gopher-heart-eyes.png")
	idea            = Atlas.AddFile("emoji/gopher-idea.png")
	insomnia        = Atlas.AddFile("emoji/gopher-insomnia.png")
	mindBlown       = Atlas.AddFile("emoji/gopher-mind-blown.png")
	neutral         = Atlas.AddFile("emoji/gopher-neutral.png")
	noPeeking       = Atlas.AddFile("emoji/gopher-no-peeking.png")
	notSureIf       = Atlas.AddFile("emoji/gopher-not-sure-if.png")
	pirate          = Atlas.AddFile("emoji/gopher-pirate.png")
	sadSweat        = Atlas.AddFile("emoji/gopher-sad-sweat.png")
	sad             = Atlas.AddFile("emoji/gopher-sad.png")
	sick            = Atlas.AddFile("emoji/gopher-sick.png")
	sleeping        = Atlas.AddFile("emoji/gopher-sleeping.png")
	sleepy          = Atlas.AddFile("emoji/gopher-sleepy.png")
	smilingBlushing = Atlas.AddFile("emoji/gopher-smiling-blushing.png")
	smilingSweat    = Atlas.AddFile("emoji/gopher-smiling-sweat.png")
	smiling         = Atlas.AddFile("emoji/gopher-smiling.png")
	thinking        = Atlas.AddFile("emoji/gopher-thinking.png")
	tired           = Atlas.AddFile("emoji/gopher-tired.png")
	tryingHard      = Atlas.AddFile("emoji/gopher-trying-hard.png")
	victorious      = Atlas.AddFile("emoji/gopher-victorious.png")
	wink            = Atlas.AddFile("emoji/gopher-wink.png")
	wondering       = Atlas.AddFile("emoji/gopher-wondering.png")
)

func main() {
	opengl.Run(run)
}

func run() {
	cfg := opengl.WindowConfig{
		Title:  "PixelUi Test",
		Bounds: pixel.R(0, 0, 1920, 1080),
	}
	win, err := opengl.NewWindow(cfg)
	if err != nil {
		log.Fatal(err)
	}

	Atlas.Pack()
	ui := pixelui.New(win, &Atlas, 0)

	for !win.Closed() {
		framestart := time.Now()
		ui.NewFrame()
		if win.JustReleased(pixel.KeyEscape) {
			win.SetClosed(true)
		}

		win.Clear(colornames.Skyblue)

		if ui.JustPressed(pixel.MouseButtonLeft) {
			fmt.Println("Left pressed")
		}

		if ui.JustReleased(pixel.MouseButtonLeft) {
			fmt.Println("Left released")
		}

		imgui.ShowDemoWindow(nil)

		imgui.Begin("Image Test")
		{
			imgui.Image(imgui.TextureID(angry.ID()), pixelui.IVec(angry.Bounds().Size()))
			imgui.Image(imgui.TextureID(atPeace.ID()), pixelui.IVec(atPeace.Bounds().Size()))
			imgui.Image(imgui.TextureID(blushing.ID()), pixelui.IVec(blushing.Bounds().Size()))
			imgui.Image(imgui.TextureID(coldSweat.ID()), pixelui.IVec(coldSweat.Bounds().Size()))
			imgui.Image(imgui.TextureID(confused.ID()), pixelui.IVec(confused.Bounds().Size()))
			imgui.Image(imgui.TextureID(cryingRiver.ID()), pixelui.IVec(cryingRiver.Bounds().Size()))
			imgui.Image(imgui.TextureID(crying.ID()), pixelui.IVec(crying.Bounds().Size()))
			imgui.Image(imgui.TextureID(dead.ID()), pixelui.IVec(dead.Bounds().Size()))
			imgui.Image(imgui.TextureID(expressionless.ID()), pixelui.IVec(expressionless.Bounds().Size()))
			imgui.Image(imgui.TextureID(facepalm.ID()), pixelui.IVec(facepalm.Bounds().Size()))
			imgui.Image(imgui.TextureID(happy.ID()), pixelui.IVec(happy.Bounds().Size()))
			imgui.Image(imgui.TextureID(heartEyes.ID()), pixelui.IVec(heartEyes.Bounds().Size()))
			imgui.Image(imgui.TextureID(idea.ID()), pixelui.IVec(idea.Bounds().Size()))
			imgui.Image(imgui.TextureID(insomnia.ID()), pixelui.IVec(insomnia.Bounds().Size()))
			imgui.Image(imgui.TextureID(mindBlown.ID()), pixelui.IVec(mindBlown.Bounds().Size()))
			imgui.Image(imgui.TextureID(neutral.ID()), pixelui.IVec(neutral.Bounds().Size()))
			imgui.Image(imgui.TextureID(noPeeking.ID()), pixelui.IVec(noPeeking.Bounds().Size()))
			imgui.Image(imgui.TextureID(notSureIf.ID()), pixelui.IVec(notSureIf.Bounds().Size()))
			imgui.Image(imgui.TextureID(pirate.ID()), pixelui.IVec(pirate.Bounds().Size()))
			imgui.Image(imgui.TextureID(sadSweat.ID()), pixelui.IVec(sadSweat.Bounds().Size()))
			imgui.Image(imgui.TextureID(sad.ID()), pixelui.IVec(sad.Bounds().Size()))
			imgui.Image(imgui.TextureID(sick.ID()), pixelui.IVec(sick.Bounds().Size()))
			imgui.Image(imgui.TextureID(sleeping.ID()), pixelui.IVec(sleeping.Bounds().Size()))
			imgui.Image(imgui.TextureID(sleepy.ID()), pixelui.IVec(sleepy.Bounds().Size()))
			imgui.Image(imgui.TextureID(smilingBlushing.ID()), pixelui.IVec(smilingBlushing.Bounds().Size()))
			imgui.Image(imgui.TextureID(smilingSweat.ID()), pixelui.IVec(smilingSweat.Bounds().Size()))
			imgui.Image(imgui.TextureID(smiling.ID()), pixelui.IVec(smiling.Bounds().Size()))
			imgui.Image(imgui.TextureID(thinking.ID()), pixelui.IVec(thinking.Bounds().Size()))
			imgui.Image(imgui.TextureID(tired.ID()), pixelui.IVec(tired.Bounds().Size()))
			imgui.Image(imgui.TextureID(tryingHard.ID()), pixelui.IVec(tryingHard.Bounds().Size()))
			imgui.Image(imgui.TextureID(victorious.ID()), pixelui.IVec(victorious.Bounds().Size()))
			imgui.Image(imgui.TextureID(wink.ID()), pixelui.IVec(wink.Bounds().Size()))
			imgui.Image(imgui.TextureID(wondering.ID()), pixelui.IVec(wondering.Bounds().Size()))
		}
		imgui.End()

		ui.Draw(win)

		win.Update()
		if dur, fDur := time.Since(framestart), time.Second/60; dur < fDur {
			time.Sleep(fDur - dur)
		}
	}
}
