package blueprint

import (
	"image/color"
)

// NewStyle will return a new element Style
func NewStyle(r Rects, p Padding, m Margin, bg, fg color.Color) (s Style) {
	s.r = r
	s.p = p
	s.bg = bg
	s.fg = fg
	return
}

// Style is an internal styling helper
type Style struct {
	// Element coordinate relative to window (absolute positioning)
	c Coords
	// Element rects (width and height)
	r Rects
	// Element padding (internal space)
	p Padding
	// Element margin (external space)
	m Margin
	// Background color
	bg color.Color
	// Foreground (used for font color)
	fg color.Color
}

// Coords will return the Style coords
func (s *Style) Coords() Coords {
	return s.c
}

// Rects will return the Style rects
func (s *Style) Rects() Rects {
	return s.r
}

// Padding will return the Style padding
func (s *Style) Padding() Padding {
	return s.p
}

// Background will return the background color for this style
func (s *Style) Background() color.Color {
	return s.bg
}

// Foreground will return the foreground for this style
func (s *Style) Foreground() color.Color {
	return s.fg
}

// Validate will validate a style
func (s *Style) Validate() (err error) {
	if s.r.Width < s.p.Left+s.p.Right || s.r.Height < s.p.Top+s.p.Bottom {
		return ErrPaddingExceedsRects
	}

	return
}
