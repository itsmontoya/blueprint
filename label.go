package blueprint

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
	"github.com/missionMeteora/journaler"
)

// NewLabel will return a new label
func NewLabel(parent Widget, str string, s Style, c Coords, f *Font) *Label {
	var l Label
	l.str = str
	l.s = s
	l.s.c = c
	l.s.r.Height = int64(f.size)

	atlas := text.NewAtlas(f.Face(), text.ASCII)
	l.f = f

	c.Y = parent.Rects().Height - l.s.r.Height
	c.Y = parent.Rects().Height - int64(float64(l.s.r.Height)/1.14)
	journaler.Debug("New label: %v %v", c, f.Size())

	l.od = c
	l.t = text.New(l.od.Vec(), atlas)
	l.t.Color = l.s.fg
	l.refresh()
	return &l
}

// Label is a label
type Label struct {
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

// Draw will draw the contents
func (l *Label) Draw(tgt pixel.Target) {
	// Draw label
	l.t.Draw(tgt, pixel.IM)
}

// Set will set a label value
func (l *Label) Set(str string) {
	l.str = str
	l.refresh()
}
