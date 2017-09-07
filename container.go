package blueprint

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/missionMeteora/journaler"
)

// NewContainer will return a new container
func NewContainer(style Style, coords Coords) *Container {
	var c Container
	c.d = imdraw.New(nil)
	c.s = style
	c.s.c = coords
	journaler.Debug("New container..: %v %v", _b, c.s.bg)
	if c.s.bg != NilColor {
		c.drawRect()
	}

	setUpdate()
	return &c
}

// Container is a standard content container
type Container struct {
	// Rectangle renderer
	d *imdraw.IMDraw
	// Container style
	s Style
	// Child widgets
	ws []Widget
}

func (c *Container) drawRect() {
	var p1, p2 Coords
	// Set point 1
	p1.X = c.s.c.X
	p1.Y = windowHeight() - c.s.c.Y
	// Set point 2
	p2.X = c.s.c.X + c.s.r.Width
	p2.Y = windowHeight() - (c.s.c.Y + c.s.r.Height)
	// Sec rectange color
	c.d.Color = c.s.bg
	// Push the first point
	c.d.Push(p1.Vec())
	c.d.Push(p2.Vec())
	c.d.Rectangle(0)
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
func (c *Container) Draw(win *pixelgl.Window) {
	// Draw container
	c.d.Draw(win)
	// Iterate through Widgets
	for _, w := range c.ws {
		w.Draw(win)
	}
}

// Push will push a widget into the container
func (c *Container) Push(w Widget) {
	c.ws = append(c.ws, w)
	setUpdate()
}

// Dot will return the next dot
func (c *Container) Dot() pixel.Vec {
	cc := c.Coords()
	cc.X += c.s.p.Left
	cc.Y += c.s.p.Top

	for _, w := range c.ws {
		cc.Y += w.Rects().Height
	}

	return cc.Vec()
}
