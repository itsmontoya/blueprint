package blueprint

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// NewCanvas will return a new Canvas
func NewCanvas(parentHeight int64, s Style) *Canvas {
	var c Canvas
	b := Bounds{
		p1: Coords{0, 0},
		p2: Coords{s.r.Width, s.r.Height},
	}

	c.Canvas = pixelgl.NewCanvas(b.PixelRect())

	topleft := Coords{
		X: (s.r.Width / 2) + s.c.X,
		Y: parentHeight - ((s.r.Height / 2) + s.c.Y),
		//Y: windowHeight() - ((s.r.Height / 2) + s.c.Y),
	}

	c.m = pixel.IM.Moved(topleft.Vec())
	return &c
}

// Canvas likes to display things
type Canvas struct {
	*pixelgl.Canvas
	m pixel.Matrix
}

// Draw will draw the canvas to a given target
func (c *Canvas) Draw(tgt pixel.Target) {
	c.Canvas.Draw(tgt, c.m)
}
