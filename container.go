package blueprint

// TODO: Make Containers thread-safe

import (
	"image/color"

	"github.com/Path94/atoms"
	"github.com/faiface/pixel"
	"github.com/missionMeteora/journaler"
)

// NewContainer will return a new container
func NewContainer(style Style, coords Coords) *Container {
	var c Container
	c.s = style
	c.s.c = coords

	c.e = NewEvents()
	c.c = NewCanvas(c.s)

	setUpdate()
	return &c
}

// Container is a standard content container
type Container struct {
	mux atoms.RWMux
	// Events embedded as the Container
	e *Events
	// Canvas represents the container in the visual form
	c *Canvas
	// Container style
	s Style
	// Child widgets
	ws []Widget
}

func (c *Container) notify(evt Event) (has bool) {
	evt.wp.X -= c.s.c.X
	evt.wp.Y -= c.s.c.Y

	journaler.Debug("Event: %v", evt.wp)
	// TODO: Remove the need for this
	topleft := Coords{
		X: (c.s.r.Width / 2) + evt.wp.X,
		Y: windowHeight() - ((c.s.r.Height / 2) + evt.wp.Y),
	}

	var dot Coords
	dot.X = evt.wp.X
	dot.Y += c.s.r.Height
	dot.Y -= evt.wp.Y

	for _, w := range c.ws {
		journaler.Debug("Checking: %v %v %v %v", w.Coords(), w.Rects(), dot, isWithinBounds(dot, w))

		if !isWithinBounds(topleft, w) {
			continue
		}

		if c, ok := w.(*Container); ok {
			if has = c.notify(evt); has {
				return
			}
		} else {
			if has = w.Events().notify(evt); has {
				return
			}
		}
	}

	return c.e.notify(evt)
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

// Events will return the container events
func (c *Container) Events() *Events {
	return c.e
}

// Draw will draw the contents
func (c *Container) Draw(tgt pixel.Target) {
	// Clear as background color
	c.c.Clear(c.s.bg)

	// Iterate through Widgets
	for _, w := range c.ws {
		w.Draw(c.c)
	}

	c.c.Draw(tgt)
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

// SetBG will change the BG color
func (c *Container) SetBG(bg color.Color) {
	c.mux.Update(func() {
		c.s.bg = bg
	})

	setUpdate()
}
