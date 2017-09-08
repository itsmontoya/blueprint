package main

import (
	"image/color"
	"os"

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

var (
	r  = blueprint.Rects{Width: 300, Height: 300}
	r2 = blueprint.Rects{Width: 96, Height: 96}
)

var (
	s1 = blueprint.NewStyle(r, blueprint.Padding{}, blueprint.Margin{}, blackC, whiteC)
	s2 = blueprint.NewStyle(r, blueprint.Padding{}, blueprint.Margin{}, whiteC, blackC)
	s3 = blueprint.NewStyle(r2, blueprint.Padding{}, blueprint.Margin{}, redC, whiteC)
	s4 = blueprint.NewStyle(r2, blueprint.Padding{}, blueprint.Margin{}, greenC, blackC)
	s5 = blueprint.NewStyle(r2, blueprint.Padding{}, blueprint.Margin{}, blueC, blackC)
)

var (
	fontRegular *blueprint.Font
)

func main() {
	var (
		a   App
		err error
	)

	if fontRegular, err = blueprint.NewFont("./assets/fonts/RobotoCondensed-Regular.ttf", 24); err != nil {
		journaler.Error("Error loading font: %v", err)
		os.Exit(1)
	}

	a.b = blueprint.New("Test app", blueprint.Rects{Width: 640, Height: 480}, blueprint.NilColor)
	if err = a.b.Run(a.Run); err != nil {
		journaler.Error("Error initializing blueprint: %v", err)
		os.Exit(1)
	}
}

// App represents our app
type App struct {
	b *blueprint.Blueprint
}

// Run will run the app!
func (a *App) Run() {
	var c *blueprint.Container
	c = blueprint.NewContainer(s1, blueprint.Coords{X: 12, Y: 12})
	a.b.Push(c)

	c = blueprint.NewContainer(s2, blueprint.Coords{X: 24, Y: 24})
	a.b.Push(c)

	c = blueprint.NewContainer(s1, blueprint.Coords{X: 36, Y: 36})
	a.b.Push(c)

	c = blueprint.NewContainer(s2, blueprint.Coords{X: 48, Y: 48})
	a.b.Push(c)

	c.Push(blueprint.NewContainer(s3, c.Dot()))
	c.Push(blueprint.NewContainer(s4, c.Dot()))
	c.Push(blueprint.NewContainer(s5, c.Dot()))
	c.Push(blueprint.NewContainer(s3, c.Dot()))

	c = blueprint.NewContainer(s1, blueprint.Coords{X: 60, Y: 60})
	a.b.Push(c)
	l := blueprint.NewLabel(c, "HELLO WORLD", s1, c.Dot(), fontRegular)
	c.Push(l)
	c.Push(blueprint.NewContainer(s3, c.Dot()))
}
