package gui

import (
	"image/png"
	"path/filepath"
	"runtime"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/atlas"
	"github.com/gopxl/pixelui/v2"
	"github.com/inkyblackness/imgui-go/v4"
)

var basepath string = getBasepath()

type MenuOptions struct {
	Padding pixel.Vec
	Height  float64

	SetAnchor   func(anchor pixel.Anchor)
	UnsetAnchor func()
	Reset       func()
}

func NewMenu(win *opengl.Window, textures *atlas.Atlas, options MenuOptions) *Menu {
	basepath = getBasepath()
	var (
		topLeft     = filepath.Join(basepath, "assets/anchor_button_top_left.png")
		top         = filepath.Join(basepath, "assets/anchor_button_top.png")
		topRight    = filepath.Join(basepath, "assets/anchor_button_top_right.png")
		left        = filepath.Join(basepath, "assets/anchor_button_left.png")
		center      = filepath.Join(basepath, "assets/anchor_button_center.png")
		right       = filepath.Join(basepath, "assets/anchor_button_right.png")
		bottomLeft  = filepath.Join(basepath, "assets/anchor_button_bottom_left.png")
		bottom      = filepath.Join(basepath, "assets/anchor_button_bottom.png")
		bottomRight = filepath.Join(basepath, "assets/anchor_button_bottom_right.png")
	)
	if options.Height == 0 {
		options.Height = 32
	}

	guiGroup := textures.MakeGroup()
	menu := &Menu{
		win:     win,
		options: options,
		anchorButtons: []*anchorButton{
			// Pixel anchors describe the direction in which the anchored item will me moved,
			// so we take the opposite in order to describe the position of the anchor itself.
			{guiGroup.AddFile(topLeft, png.Decode), pixel.TopLeft.Opposite()},
			{guiGroup.AddFile(top, png.Decode), pixel.Top.Opposite()},
			{guiGroup.AddFile(topRight, png.Decode), pixel.TopRight.Opposite()},
			{guiGroup.AddFile(left, png.Decode), pixel.Left.Opposite()},
			{guiGroup.AddFile(center, png.Decode), pixel.Center},
			{guiGroup.AddFile(right, png.Decode), pixel.Right.Opposite()},
			{guiGroup.AddFile(bottomLeft, png.Decode), pixel.BottomLeft.Opposite()},
			{guiGroup.AddFile(bottom, png.Decode), pixel.Bottom.Opposite()},
			{guiGroup.AddFile(bottomRight, png.Decode), pixel.BottomRight.Opposite()},
		},
	}
	textures.Pack()
	return menu
}

// Menu has controls anchoring and reseting text
type Menu struct {
	win     *opengl.Window
	options MenuOptions

	anchorButtons []*anchorButton
	selected      *anchorButton
}

// Update is called every frame to redraw UI elements in the window
func (m *Menu) Update() {
	imgui.PushStyleVarVec2(imgui.StyleVarWindowPadding, pixelui.IVec(m.options.Padding))
	imgui.SetNextWindowPos(pixelui.IVec(pixel.V(0, 0)))
	imgui.SetNextWindowSize(pixelui.IVec(m.size()))
	imgui.BeginV("Menu", nil, imgui.WindowFlagsNoDecoration)
	{
		m.anchorDropdown()
		m.resetButton()
	}
	imgui.End()
	imgui.PopStyleVarV(1)
}

func (m *Menu) size() pixel.Vec {
	return pixel.V(m.win.Bounds().Size().X, m.options.Height)
}

func (m *Menu) buttonHeight() float32 {
	return float32(m.options.Height - 2*m.options.Padding.Y)
}

// anchorDropdown manages the anchor button popup window
func (m *Menu) anchorDropdown() {
	size := imgui.Vec2{X: 64, Y: m.buttonHeight()}
	imgui.ButtonV("Anchor", size)
	imgui.SetNextWindowPos(imgui.Vec2{X: float32(m.options.Padding.X), Y: float32(m.options.Padding.Y) + size.Y + 4})
	imgui.OpenPopupOnItemClickV("Anchors", 0)
	if imgui.BeginPopup("Anchors") {
		for i, anchorButton := range m.anchorButtons {
			if i%3 != 0 {
				imgui.SameLineV(0, 4)
			}
			active := anchorButton == m.selected
			if anchorButton.add(active) {
				if active {
					m.selected = nil
					m.options.UnsetAnchor()
				} else {
					m.selected = anchorButton
					m.options.SetAnchor(anchorButton.anchor)
				}
			}
		}
		imgui.EndPopup()
	}
}

func (m *Menu) resetButton() {
	size := imgui.Vec2{X: 48, Y: m.buttonHeight()}
	imgui.SameLine()
	imgui.PushStyleColor(imgui.StyleColorButton, pixelui.ColorA(255, 0, 0, 200))
	imgui.PushStyleColor(imgui.StyleColorButtonHovered, pixelui.ColorA(220, 0, 0, 200))
	imgui.PushStyleColor(imgui.StyleColorButtonActive, pixelui.ColorA(175, 0, 0, 200))
	if imgui.ButtonV("Reset", size) {
		m.selected = nil
		m.options.Reset()
	}
	imgui.PopStyleColorV(3)
}

type anchorButton struct {
	textureId atlas.TextureId
	anchor    pixel.Anchor
}

func (ab *anchorButton) add(active bool) bool {
	var color imgui.Vec4
	if active {
		color = pixelui.ColorA(255, 255, 255, 100)
	} else {
		color = pixelui.Color(255, 255, 255)
	}

	if imgui.ImageButtonV(
		imgui.TextureID(ab.textureId.ID()),
		pixelui.IVec(ab.textureId.Bounds().Size()),
		imgui.Vec2{X: 0, Y: 0},
		imgui.Vec2{X: 1, Y: 1},
		0,
		imgui.Vec4{X: 0, Y: 0, Z: 0, W: 0},
		color,
	) {
		return true
	}
	return false
}

func getBasepath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.ToSlash(filepath.Dir(filepath.Dir(b)))
}
