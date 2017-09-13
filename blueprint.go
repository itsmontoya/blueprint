package blueprint

import (
	"image/color"
	"sync"

	"github.com/Path94/atoms"
	"github.com/faiface/pixel/pixelgl"
	"github.com/missionMeteora/journaler"
	"github.com/missionMeteora/toolkit/errors"
)

const (
	// ErrNotInitialized is returned when an action is performed before blueprint has been initialized
	ErrNotInitialized = errors.Error("blueprint has not yet been initialized")
	// ErrPaddingExceedsRects is returned when an element padding is greater than the element rects
	ErrPaddingExceedsRects = errors.Error("element padding exceeds it's rects")
)

var (
	// NilColor represents a nil color value
	NilColor = color.RGBA{255, 105, 180, 1}
	// TransparentColor represents a transparent color
	TransparentColor = color.RGBA{0, 0, 0, 0}
)

var (
	_b *Blueprint
)

// New will return a new instance of Blueprint
func New(title string, rects Rects, bg color.Color) *Blueprint {
	var b Blueprint
	b.out = journaler.New(title)

	b.title = title
	b.rects = rects
	b.bg = bg
	b.setUpdate()

	_b = &b
	return &b
}

// Blueprint manages an application GUI
type Blueprint struct {
	mux atoms.RWMux

	title string
	rects Rects

	win *pixelgl.Window
	out *journaler.Journaler

	ws []Widget
	bg color.Color

	hovering Widget

	// Update state
	update atoms.Bool
	// Running wait group
	wg sync.WaitGroup
	// Closed state
	closed atoms.Bool
}

func (b *Blueprint) setUpdate() {
	b.update.Set(true)
}

func (b *Blueprint) render() {
	b.win.Clear(b.bg)
	for _, w := range b.ws {
		w.Draw(b.win)
	}
}

// Push will push a widget onto the Blueprint
func (b *Blueprint) Push(w Widget) {
	b.mux.Update(func() {
		b.ws = append(b.ws, w)
		b.setUpdate()
	})
}

// Run will begin the render loop, function will end when window or the instance of blueprint has been closed
func (b *Blueprint) Run(fn func()) (err error) {
	b.wg.Add(1)

	pixelgl.Run(func() {
		if b.win, err = pixelgl.NewWindow(getCfg(b.title, b.rects.Width, b.rects.Height)); err != nil {
			return
		}

		go fn()

		for !b.win.Closed() {
			b.mux.Update(func() {
				b.checkHovering()
				b.checkMouseDown()

				if b.update.Set(false) {
					b.render()
				}

				b.win.Update()
			})
		}
	})

	b.wg.Done()
	return
}

// Rects will return the Blueprint rects
func (b *Blueprint) Rects() (r Rects) {
	b.mux.Read(func() {
		r = b.rects
	})

	return
}

func (b *Blueprint) handleMouseLeave(evt Event) {
	if b.hovering == nil {
		return
	}

	if c, ok := b.hovering.(*Container); ok {
		c.notify(evt)
	}

	b.hovering.Events().notify(evt)
	b.hovering = nil
}

func (b *Blueprint) checkHovering() {
	var enter Event
	enter.et = EventMouseEnter
	enter.wp = newCoordsFromVec(b.win.MousePosition())
	enter.wp.Y = windowHeight() - enter.wp.Y

	leave := enter
	leave.et = EventMouseLeave

	for i := len(b.ws) - 1; i > -1; i-- {
		w := b.ws[i]
		if !isWithinBounds(enter.wp, w) {
			continue
		}

		updating := w != b.hovering
		if updating {
			b.handleMouseLeave(leave)
		}

		// We notify containers on every instance of hover so a check can be made for it's children
		if c, ok := w.(*Container); ok {
			c.notify(enter)
		}

		if updating {
			// We notify direct widgets once on hover
			w.Events().notify(enter)
		}

		b.hovering = w
		return
	}

	// If we made it to the end, we have no matches
	b.handleMouseLeave(leave)
}

func (b *Blueprint) checkMouseDown() {
	if !b.win.JustPressed(pixelgl.MouseButton1) {
		return
	}

	var evt Event
	evt.et = EventMouseDown
	evt.wp = newCoordsFromVec(b.win.MousePosition())
	evt.wp.Y = windowHeight() - evt.wp.Y

	for i := len(b.ws) - 1; i > -1; i-- {
		w := b.ws[i]
		if !isWithinBounds(evt.wp, w) {
			continue
		}

		if c, ok := w.(*Container); ok {
			if c.notify(evt) {
				break
			}
		} else {
			if w.Events().notify(evt) {
				break
			}
		}

		break
	}
}

func isWithinBounds(pos Coords, w Widget) (within bool) {
	c := w.Coords()
	if pos.X < c.X || pos.Y < c.Y {
		return
	}

	r := w.Rects()
	if pos.X > c.X+r.Width || pos.Y > c.Y+r.Height {
		return
	}

	return true
}

// Close will blue an instance of Blueprint
func (b *Blueprint) Close() (err error) {
	if !b.closed.Set(true) {
		return errors.ErrIsClosed
	}

	b.win.SetClosed(true)
	b.wg.Wait()
	return
}

func setUpdate() {
	if _b == nil {
		return
	}

	_b.setUpdate()
}

func windowWidth() int64 {
	if _b == nil {
		return 0
	}

	return _b.rects.Width
}

func windowHeight() int64 {
	if _b == nil {
		return 0
	}

	return _b.rects.Height
}
