package main

import (
	"fmt"
	"image"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"

	"golang.org/x/image/colornames"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/math/fixed"
)

var (
	defaultFont font.Face
)

func init() {
	f, err := freetype.ParseFont(gomono.TTF)
	if err != nil {
		panic(err)
	}
	defaultFont = truetype.NewFace(f, &truetype.Options{
		Size: 14,
	})
	if err != nil {
		panic(err)
	}
}

func toImageRect(r pixel.Rect) image.Rectangle {
	return image.Rect(int(r.Min.X), int(r.Min.Y), int(r.Max.X), int(r.Max.Y))
}

type (
	Lines struct {
		baseNode
		text map[int]string
		min  int
		max  int

		nextFrame chan image.Image

		changed     bool
		initialized bool
		content     pixel.Picture
		frameRect   pixel.Rect
		center      pixel.Vec
		grid        gridInfo
	}

	// contains information about the available character grid
	// mostly used to know how many rows/cols are available
	gridInfo struct {
		rows int
		cols int
	}
)

func NewLines() *Lines {
	return &Lines{
		baseNode: newBaseNode(),
	}
}

func (l *Lines) Render(rc *RenderContext) {
	select {
	case im := <-l.nextFrame:
		l.content = pixel.PictureDataFromImage(im)
	default:
	}
	if l.content == nil {
		return
	}

	sp := pixel.NewSprite(l.content, l.content.Bounds())
	sp.Draw(rc.Target, pixel.IM.Moved(l.center))
}

func (l *Lines) Update(win *pixelgl.Window) error {
	if !l.initialized {
		l.init()
	}
	l.setText(1, time.Now().Format(time.RFC3339Nano))
	l.setText(2, "second line")
	if win.Bounds().Size() != l.frameRect.Size() {
		sz := win.Bounds().Size()
		l.frameRect = pixel.R(0, 0, sz.X, sz.Y)
		l.grid.recomputeSize(defaultFont, l.frameRect)
		for i := 5; i <= l.grid.rows; i++ {
			l.setText(i, "...")
		}
	}
	l.setText(3, fmt.Sprintf("Grid size: %v x %v", l.grid.cols, l.grid.rows))
	l.setText(4, fmt.Sprintf(fmt.Sprintf("%%0%vd", l.grid.cols), l.grid.cols))
	l.renderFrame()
	l.center = win.Bounds().Center()
	return nil
}

func (l *Lines) setText(line int, text string) {
	l.text[line] = text
	l.expandToFit(line)
}

func (l *Lines) expandToFit(nl int) {
	if l.min > nl {
		l.min = nl
	}
	if l.max <= nl {
		l.max = nl + 1
	}
}

func (l *Lines) renderFrame() {
	d := font.Drawer{
		Dst:  image.NewRGBA(toImageRect(l.frameRect)),
		Src:  image.NewUniform(colornames.Black),
		Dot:  fixed.P(0, defaultFont.Metrics().Height.Ceil()),
		Face: defaultFont,
	}
	for i := l.min; i < l.max; i++ {
		d.DrawString(l.text[i])
		d.Dot = fixed.P(0, d.Dot.Y.Ceil()+defaultFont.Metrics().Height.Ceil())
	}
	l.nextFrame <- d.Dst
}

func (l *Lines) init() {
	if l.text == nil {
		l.text = make(map[int]string)
	}
	l.min, l.max = 1, 0
	for k := range l.text {
		if k <= 0 {
			delete(l.text, k)
			continue
		}
		l.expandToFit(k)
	}
	l.nextFrame = make(chan image.Image, 1)
}

func (gi *gridInfo) recomputeSize(f font.Face, r pixel.Rect) {
	mglyph, _ := f.GlyphAdvance('M')
	glyphHeight := f.Metrics().Height
	gi.cols = int(r.W()) / mglyph.Ceil()
	gi.rows = int(r.H()) / glyphHeight.Ceil()
	println("mglyph: ", mglyph.String(), "round: ", mglyph.Round(), "ceil: ", mglyph.Ceil())
}
