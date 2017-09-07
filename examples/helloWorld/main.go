package main

import (
	"image/color"
	"os"

	"github.com/itsmontoya/blueprint"
	"github.com/missionMeteora/journaler"
)

var (
	blackC = color.RGBA{0, 0, 0, 255}
	whiteC = color.RGBA{255, 255, 255, 255}
)

func main() {
	var (
		b = blueprint.New("Test app", blueprint.Rects{Width: 640, Height: 480}, blueprint.NilColor)
		r = blueprint.Rects{Width: 300, Height: 250}

		c *blueprint.Container

		s1, s2 blueprint.Style
		err    error
	)

	if s1, err = blueprint.NewStyle(r, blueprint.Padding{}, blueprint.Margin{}, blackC, whiteC); err != nil {
		journaler.Error("Error getting new style: %v", err)
		os.Exit(1)
	}

	if s2, err = blueprint.NewStyle(r, blueprint.Padding{}, blueprint.Margin{}, whiteC, blackC); err != nil {
		journaler.Error("Error getting new style: %v", err)
		os.Exit(1)
	}

	c = blueprint.NewContainer(s1, blueprint.Coords{X: 12, Y: 12})
	b.Push(c)

	c = blueprint.NewContainer(s2, blueprint.Coords{X: 24, Y: 24})
	b.Push(c)

	c = blueprint.NewContainer(s1, blueprint.Coords{X: 36, Y: 36})
	b.Push(c)

	c = blueprint.NewContainer(s2, blueprint.Coords{X: 48, Y: 48})
	b.Push(c)

	if err = b.Run(); err != nil {
		journaler.Error("Error initializing blueprint: %v", err)
		os.Exit(1)
	}
}
