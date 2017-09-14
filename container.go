package blueprint

// TODO: Make Containers thread-safe

import (
	"image/color"

	"github.com/Path94/atoms"
	"github.com/faiface/pixel"
)

const (
	matchNone uint8 = iota
	matchOK
	matchBreak
)

var nilWidget Widget

// NewContainer will return a new container
func NewContainer(p Parent, s Style, c Coords) *Container {
	var cnt Container
	cnt.s = s
	cnt.s.c = c

	cnt.p = p
	cnt.e = NewEvents()
	cnt.c = NewCanvas(p.Rects().Height, cnt.s)
	cnt.hovering = nilWidget
	cnt.SetToUpdate()
	return &cnt
}

// Container is a standard content container
type Container struct {
	mux atoms.RWMux
	// Container parent
	p Parent
	// Events embedded as the Container
	e *Events
	// Canvas represents the container in the visual form
	c *Canvas
	// Container style
	s Style

	hovering Widget

	// Child widgets
	ws []Widget
}

func (c *Container) handleMouseEnter(evt Event, w Widget) (has bool) {
	if c.hovering != w {
		c.handleMouseLeave(evt)
	}

	if cc, ok := w.(*Container); ok {
		has = cc.notify(evt)
	} else if w != c.hovering {
		has = w.Events().notify(evt)
	}

	c.hovering = w
	return
}

func (c *Container) handleMouseLeave(evt Event) (has bool) {
	if c.hovering == nil {
		return
	}

	evt.et = EventMouseLeave
	if hc, ok := c.hovering.(*Container); ok {
		has = hc.notify(evt)
	} else {
		has = c.hovering.Events().notify(evt)
	}

	c.hovering = nil
	return
}

func (c *Container) handleMouseDown(evt Event, w Widget) (has bool) {
	if cc, ok := w.(*Container); ok {
		return cc.notify(evt)
	}

	return w.Events().notify(evt)
}

func (c *Container) notify(evt Event) (has bool) {
	evt.rp.X = evt.wp.X - c.s.c.X
	evt.rp.Y = evt.wp.Y - c.s.c.Y

	switch evt.et {
	case EventMouseLeave:
		// No need to iterate through widgets if out container isn't even selected!
		c.handleMouseLeave(evt)
	}

	for _, w := range c.ws {
		if !isWithinBounds(evt.rp, w) {
			continue
		}

		has = true

		switch evt.et {
		case EventMouseEnter:
			c.handleMouseEnter(evt, w)
		case EventMouseDown:
			c.handleMouseDown(evt, w)
		}

		return
	}

	switch evt.et {
	case EventMouseEnter:
		c.handleMouseLeave(evt)
	}

	return
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

func (c *Container) refresh() {
	// Clear as background color
	c.c.Clear(c.s.bg)

	// Iterate through Widgets
	for _, w := range c.ws {
		w.Draw(c.c)
	}
}

// Draw will draw the contents
func (c *Container) Draw(tgt pixel.Target) {
	c.c.Draw(tgt)
}

// Push will push a widget into the container
func (c *Container) Push(ws ...Widget) {
	c.ws = append(c.ws, ws...)
	c.SetToUpdate()
}

// Dot will return the next dot
func (c *Container) Dot() (dot Coords) {
	dot.X = c.s.p.Left
	dot.Y = c.s.p.Top

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

	c.SetToUpdate()
}

// SetToUpdate will redraw the container (and it's children) and signal the renderer for an repaint
func (c *Container) SetToUpdate() {
	c.refresh()
	c.p.SetToUpdate()
}
