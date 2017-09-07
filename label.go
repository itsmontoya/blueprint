package blueprint

/*
import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
)

func newLabel(str string, s Style, c Coords) *Label {
	var l Label
	l.str = str
	l.w = &widget{
		coords:  coords,
		padding: padding,
	}

	l.w.rects.Height = int64(f.size)

	atlas := text.NewAtlas(f.Face(), text.ASCII)
	l.f = f
	l.od = pixel.V(l.w.coords.X, windowHeight-l.w.coords.Y)
	l.t = text.New(l.od, atlas)
	l.t.Color = c
	l.refresh()
	return &l
}

// Label is a label
type Label struct {
	f *Font
	t *text.Text
	// original dot
	od pixel.Vec

	s   Style
	str string
}

func (l *Label) refresh() {
	l.t.Clear()
	l.t.Dot = l.od
	l.t.WriteString(l.str)
	l.w.setToUpdate()

}

// Coords will return the box coords
func (l *Label) Coords() pixel.Vec {
	return l.w.coords
}

// Rects will return the box rects
func (l *Label) Rects() Rects {
	return l.w.rects
}

// Padding will return the box padding
func (l *Label) Padding() Padding {
	return l.w.padding
}

// Updated will return the box update state
func (l *Label) Updated() bool {
	return l.w.updated.Set(false)
}

// Draw will draw the contents
func (l *Label) Draw(win *pixelgl.Window) {
	// Draw label
	l.t.Draw(win, pixel.IM)
}

// Set will set a label value
func (l *Label) Set(str string) {
	l.str = str
	l.refresh()
}
*/
