package blueprint

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// NewContainer will return a new container
func NewContainer(style Style, coords Coords) *Container {
	var c Container
	b := Bounds{
		p1: Coords{0, 0},
		p2: Coords{style.r.Width, style.r.Height},
	}

	c.c = pixelgl.NewCanvas(b.PixelRect())
	c.s = style
	c.s.c = coords

	setUpdate()
	return &c
}

// Container is a standard content container
type Container struct {
	c *pixelgl.Canvas
	// Container style
	s Style
	// Child widgets
	ws []Widget
}

// Coords will return the container coords
func (c *Container) Coords() Coords {
	return c.s.c
}

// Rects will return the container rects
func (c *Container) Rects() Rects {
	return c.s.r
}

// Padding will return the container padding
func (c *Container) Padding() Padding {
	return c.s.p
}

// Margin will return the container margin
func (c *Container) Margin() Margin {
	return c.s.m
}

// Draw will draw the contents
func (c *Container) Draw(tgt pixel.Target) {
	// Clear as background color
	c.c.Clear(c.s.bg)

	// Iterate through Widgets
	for _, w := range c.ws {
		w.Draw(c.c)
	}

	topleft := Coords{
		X: (c.s.r.Width / 2) + c.s.c.X,
		Y: windowHeight() - ((c.s.r.Height / 2) + c.s.c.Y),
	}

	c.c.Draw(tgt, pixel.IM.Moved(topleft.Vec()))
}

// Push will push a widget into the container
func (c *Container) Push(w Widget) {
	c.ws = append(c.ws, w)
	setUpdate()
}

// Dot will return the next dot
func (c *Container) Dot() (dot Coords) {
	dot.X += c.s.p.Left
	dot.Y += windowHeight() - c.s.r.Height
	dot.Y += c.s.p.Top

	for _, w := range c.ws {
		dot.Y += w.Rects().Height
	}

	return
}
