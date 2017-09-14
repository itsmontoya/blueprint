package blueprint

import (
	"github.com/faiface/pixel"
)

// Widget is a widget interface
type Widget interface {
	Coords() Coords
	Rects() Rects
	Padding() Padding
	Margin() Margin
	Draw(pixel.Target)
	Events() *Events

	SetToUpdate()
}
