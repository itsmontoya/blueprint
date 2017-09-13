package blueprint

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
)

// TODO: Figure out why the hell this is needed to align properly
const magicFontNumber = .74

// NewLabel will return a new label
func NewLabel(p Parent, str string, s Style, c Coords, f *Font) *Label {
	var l Label
	l.str = str
	l.s = s
	l.s.c = c
	if sz := int64(f.size); l.s.r.Height < sz {
		l.s.r.Height = sz
	}

	l.e = NewEvents()
	l.c = NewCanvas(p.Rects().Height, l.s)

	atlas := text.NewAtlas(f.Face(), text.ASCII)
	l.f = f

	l.od = Coords{X: l.s.p.Left, Y: 0}
	// Bring up cursor dot to vertical center (Note: dot indicates the BOTTOM of the chars)
	l.od.Y = (l.s.r.Height / 2)
	// Bring down the dot to compensate for font size
	l.od.Y -= int64(f.size*magicFontNumber) / 2

	l.t = text.New(l.od.Vec(), atlas)
	l.t.Color = l.s.fg
	l.refresh()
	return &l
}

// Label is a label
type Label struct {
	e *Events
	c *Canvas

	f *Font
	t *text.Text
	// original dot
	od Coords

	s   Style
	str string
}

func (l *Label) refresh() {
	l.t.Clear()
	l.t.Dot = l.od.Vec()
	l.t.WriteString(l.str)

	// Clear as background color
	l.c.Clear(l.s.bg)
	// Draw label to canvas
	l.t.Draw(l.c, pixel.IM)

	setUpdate()
}

// Coords will return the label coords
func (l *Label) Coords() Coords {
	return l.s.c
}

// Rects will return the label rects
func (l *Label) Rects() Rects {
	return l.s.r
}

// Padding will return the label padding
func (l *Label) Padding() Padding {
	return l.s.p
}

// Margin will return the label margin
func (l *Label) Margin() Margin {
	return l.s.m
}

// Events will return the label events
func (l *Label) Events() *Events {
	return l.e
}

// Draw will draw the contents
func (l *Label) Draw(tgt pixel.Target) {
	l.c.Draw(tgt)
}

// Set will set a label value
func (l *Label) Set(str string) {
	l.str = str
	l.refresh()
}

// SetBG will set the background color
func (l *Label) SetBG(bg color.Color) {
	l.s.bg = bg
	l.refresh()

}
