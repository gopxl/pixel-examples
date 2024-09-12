package main

import (
	"image/color"
	"time"
	"unicode"

	"github.com/golang/freetype/truetype"
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/atlas"
	"github.com/gopxl/pixel/v2/ext/imdraw"
	"github.com/gopxl/pixel/v2/ext/text"
	"github.com/gopxl/pixelui/v2"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/gopxl/pixel-examples/ext/text/gui"
)

var (
	textures  atlas.Atlas
	fontAtlas = text.NewAtlas(
		ttfFromBytesMust(goregular.TTF, 42),
		text.ASCII,
		text.RangeTable(unicode.Latin),
	)
)

func run() {
	title := "Text example"
	cfg := opengl.WindowConfig{
		Title:  title,
		Bounds: pixel.R(0, 0, 1024, 768),
	}
	win, err := opengl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	origin := win.Bounds().Center()
	txt := text.New(origin, fontAtlas)
	txtBox := newTextBox(txt)
	txtBox.txt.LineHeight *= 1.5

	// Create UI controls
	menu := gui.NewMenu(win, &textures, gui.MenuOptions{
		Height:      32.0,
		Padding:     pixel.V(4, 4),
		SetAnchor:   txtBox.setAnchor,
		UnsetAnchor: txtBox.unsetAnchor,
		Reset:       txtBox.reset,
	})
	ui := pixelui.New(win, &textures, 0)

	for !win.Closed() {
		win.Clear(colornames.Skyblue)

		// Draw UI
		ui.NewFrame()
		menu.Update()
		ui.Draw(win)

		// Draw text box
		txtBox.process(win)
		txtBox.draw(win)

		win.Update()
	}
}

func newCursor(color color.Color, pos pixel.Vec, thickness, lineHeight float64) *cursor {
	imd := imdraw.New(nil)
	imd.Color = color
	imd.Push(pixel.ZV, pixel.V(0, lineHeight))
	imd.Line(thickness)
	return &cursor{
		imd:     imd,
		pos:     pos,
		offset:  pixel.V(6, -2),
		blink:   time.NewTicker(time.Second),
		visible: true,
	}
}

// cursor draws a blinking text cursor at the current Dot position of the text
type cursor struct {
	imd     *imdraw.IMDraw
	pos     pixel.Vec
	blink   *time.Ticker
	visible bool
	offset  pixel.Vec
}

func (c *cursor) setPos(pos pixel.Vec) {
	if !c.pos.Eq(pos) {
		c.pos = pos
		c.blink.Reset(time.Second)
		c.visible = true
	}
}

func (c *cursor) draw(win *opengl.Window) {
	select {
	case <-c.blink.C:
		c.visible = !c.visible
	default:
	}
	if c.visible {
		win.SetMatrix(pixel.IM.Moved(c.pos.Add(c.offset)))
		c.imd.Draw(win)
		win.SetMatrix(pixel.IM)
	}
}

func newTextBox(txt *text.Text) *textBox {
	return &textBox{
		txt:    txt,
		cursor: newCursor(pixel.RGB(1, 1, 1), txt.Orig, 2, txt.LineHeight),
		imd:    imdraw.New(nil),
		dirty:  true,
	}
}

// textBox manages updating and anchoring text
type textBox struct {
	txt    *text.Text
	cursor *cursor
	imd    *imdraw.IMDraw
	anchor pixel.Anchor

	s     string
	dirty bool
}

func (tb *textBox) process(win *opengl.Window) {
	typed := win.Typed()
	if typed != "" {
		tb.writeString(typed)
	}
	if win.JustPressed(pixel.KeyTab) || win.Repeated(pixel.KeyTab) {
		tb.writeRune('\t')
	}
	if win.JustPressed(pixel.KeyEnter) ||
		win.Repeated(pixel.KeyEnter) ||
		win.JustPressed(pixel.KeyKPEnter) ||
		win.Repeated(pixel.KeyKPEnter) {
		tb.writeRune('\n')
	}
	if win.JustPressed(pixel.KeyBackspace) || win.Repeated(pixel.KeyBackspace) {
		tb.delete()
	}

	if tb.dirty {
		tb.cursor.setPos(tb.txt.AnchoredDot())
	}
}

func (tb *textBox) writeString(s string) (n int, err error) {
	tb.s += s
	tb.dirty = true
	return tb.txt.WriteString(s)
}

func (tb *textBox) writeRune(r rune) (n int, err error) {
	tb.s += string(r)
	tb.dirty = true
	return tb.txt.WriteRune(r)
}

func (tb *textBox) reset() {
	tb.dirty = true
	tb.s = ""
	tb.txt.Clear()
	tb.txt.Unaligned()
}

func (tb *textBox) unsetAnchor() {
	tb.txt.Unaligned()
	tb.dirty = true
}

func (tb *textBox) delete() {
	if len(tb.s) == 0 {
		return
	}
	tb.dirty = true
	tb.s = tb.s[:len(tb.s)-1]
	tb.txt.Clear()
	tb.txt.WriteString(tb.s)
}

func (tb *textBox) draw(win *opengl.Window) {
	if tb.dirty {
		tb.redraw()
	}
	tb.txt.Draw(win, pixel.IM)
	tb.imd.Draw(win)
	tb.cursor.draw(win)
}

// redraw updates imdraw elements when the text has changed
func (tb *textBox) redraw() {
	tb.imd.Clear()

	// Bounds
	tb.imd.Color = colornames.Red
	bounds := tb.txt.AnchoredBounds()
	tb.imd.Push(bounds.Min, bounds.Max)
	tb.imd.Rectangle(2)

	// Origin
	tb.imd.Color = colornames.Blue
	tb.imd.Push(tb.txt.Orig)
	tb.imd.Circle(3, 1)

	// Dot
	tb.imd.Color = colornames.Green
	tb.imd.Push(tb.txt.AnchoredDot())
	tb.imd.Circle(3, 1)

	tb.dirty = false
}

func (tb *textBox) setAnchor(anchor pixel.Anchor) {
	tb.txt.AlignedTo(anchor)
	tb.anchor = anchor
	tb.dirty = true
}

func ttfFromBytesMust(b []byte, size float64) font.Face {
	ttf, err := truetype.Parse(b)
	if err != nil {
		panic(err)
	}
	return truetype.NewFace(ttf, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	})
}

func main() {
	opengl.Run(run)
}
