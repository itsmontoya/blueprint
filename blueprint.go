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
	ws  []Widget
	bg  color.Color
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
			if b.win.JustPressed(pixelgl.MouseButton1) {
				journaler.Debug("Click: %v", b.win.MousePosition())
			}

			if b.update.Set(false) {
				b.mux.Update(b.render)
			}

			b.win.Update()
		}
	})

	b.wg.Done()
	return
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
