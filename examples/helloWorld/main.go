package main

import (
	"image/color"
	"os"
	//	"time"

	"github.com/itsmontoya/blueprint"
	"github.com/missionMeteora/journaler"
)

var (
	blackC = color.RGBA{0, 0, 0, 255}
	grayC  = color.RGBA{100, 100, 100, 255}
	redC   = color.RGBA{255, 0, 0, 255}
	greenC = color.RGBA{0, 255, 0, 255}
	blueC  = color.RGBA{0, 0, 255, 255}
	whiteC = color.RGBA{255, 255, 255, 255}
)

var (
	r  = blueprint.Rects{Width: 300, Height: 300}
	r2 = blueprint.Rects{Width: 96, Height: 96}
	r3 = blueprint.Rects{Width: 300, Height: 64}
)

var (
	s1 = blueprint.NewStyle(r, blueprint.Padding{}, blueprint.Margin{}, blackC, whiteC)
	s2 = blueprint.NewStyle(r, blueprint.Padding{}, blueprint.Margin{}, whiteC, blackC)
	s3 = blueprint.NewStyle(r2, blueprint.Padding{}, blueprint.Margin{}, redC, whiteC)
	s4 = blueprint.NewStyle(r2, blueprint.Padding{}, blueprint.Margin{}, greenC, blackC)
	s5 = blueprint.NewStyle(r2, blueprint.Padding{}, blueprint.Margin{}, blueC, blackC)
	s6 = blueprint.NewStyle(r3, blueprint.Padding{Left: 12}, blueprint.Margin{}, grayC, whiteC)
)

var (
	fontRegular *blueprint.Font
)

func main() {
	var (
		a   App
		err error
	)

	if fontRegular, err = blueprint.NewFont("./assets/fonts/RobotoCondensed-Regular.ttf", 32); err != nil {
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
	c := blueprint.NewContainer(s1, blueprint.Coords{X: 60, Y: 60})
	a.b.Push(c)
	l := blueprint.NewLabel(c, "Hello World.", s6, c.Dot(), fontRegular)
	c.Push(l)
	c.Events().Subscribe(blueprint.EventMouseDown, func(evt blueprint.Event) {
		journaler.Success("Oh hai.")
	})

	l.Events().Subscribe(blueprint.EventMouseDown, func(evt blueprint.Event) {
		journaler.Success("Headah!")
	})
	/*
		go func() {
			redBox := blueprint.NewContainer(s3, c.Dot())
			c.Push(redBox)
			for {
				time.Sleep(time.Second)
				redBox.SetBG(blackC)

				time.Sleep(time.Second)
				redBox.SetBG(grayC)

				time.Sleep(time.Second)
				redBox.SetBG(whiteC)
			}
		}()

		for {
			time.Sleep(time.Second)
			c.SetBG(blackC)

			time.Sleep(time.Second)
			c.SetBG(greenC)

			time.Sleep(time.Second)
			c.SetBG(blueC)
		}

	*/
}
