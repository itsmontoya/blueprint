package main

import (
	"image/color"
	"os"
	"time"

	"github.com/itsmontoya/blueprint"
	"github.com/missionMeteora/journaler"
)

var (
	blackC = color.RGBA{0, 0, 0, 255}
	redC   = color.RGBA{255, 0, 0, 255}
	greenC = color.RGBA{0, 255, 0, 255}
	blueC  = color.RGBA{0, 0, 255, 255}
	whiteC = color.RGBA{255, 255, 255, 255}
)

func main() {
	var (
		b  = blueprint.New("Test app", blueprint.Rects{Width: 640, Height: 480}, blueprint.NilColor)
		r  = blueprint.Rects{Width: 300, Height: 300}
		r2 = blueprint.Rects{Width: 100, Height: 100}

		c *blueprint.Container

		s1, s2, s3, s4, s5 blueprint.Style

		err error
	)

	if s1, err = blueprint.NewStyle(r, blueprint.Padding{}, blueprint.Margin{}, blackC, whiteC); err != nil {
		journaler.Error("Error getting new style: %v", err)
		os.Exit(1)
	}

	if s2, err = blueprint.NewStyle(r, blueprint.Padding{}, blueprint.Margin{}, whiteC, blackC); err != nil {
		journaler.Error("Error getting new style: %v", err)
		os.Exit(1)
	}

	if s3, err = blueprint.NewStyle(r2, blueprint.Padding{}, blueprint.Margin{}, redC, whiteC); err != nil {
		journaler.Error("Error getting new style: %v", err)
		os.Exit(1)
	}

	if s4, err = blueprint.NewStyle(r2, blueprint.Padding{}, blueprint.Margin{}, greenC, blackC); err != nil {
		journaler.Error("Error getting new style: %v", err)
		os.Exit(1)
	}

	if s5, err = blueprint.NewStyle(r2, blueprint.Padding{}, blueprint.Margin{}, blueC, blackC); err != nil {
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

	go func() {
		time.Sleep(time.Second)
		c.Push(blueprint.NewContainer(s3, c.Dot()))
		time.Sleep(time.Second)
		c.Push(blueprint.NewContainer(s4, c.Dot()))
		time.Sleep(time.Second)
		c.Push(blueprint.NewContainer(s5, c.Dot()))
	}()
	if err = b.Run(); err != nil {
		journaler.Error("Error initializing blueprint: %v", err)
		os.Exit(1)
	}
}
