package blueprint

import (
	"github.com/faiface/pixel/pixelgl"
)

// Widget is a widget interface
type Widget interface {
	Coords() Coords
	Rects() Rects
	Padding() Padding
	Margin() Margin
	Draw(win *pixelgl.Window)
}
