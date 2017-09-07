package blueprint

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// Rects represents rect values
type Rects struct {
	Width  int64
	Height int64
}

// Padding represents padding values
type Padding struct {
	Top    int64
	Right  int64
	Bottom int64
	Left   int64
}

// Margin represents margin values
type Margin struct {
	Top    int64
	Right  int64
	Bottom int64
	Left   int64
}

// Coords represents coordinate values
type Coords struct {
	X int64
	Y int64
}

// Vec will return a Vec representation of a Coord
func (c *Coords) Vec() pixel.Vec {
	return pixel.V(float64(c.X), float64(c.Y))
}

func getCfg(title string, w, h int64) (cfg pixelgl.WindowConfig) {
	cfg.Title = title
	cfg.Bounds = pixel.R(0, 0, float64(w), float64(h))
	cfg.VSync = true
	return
}
