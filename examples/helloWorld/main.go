package main

import (
	"github.com/itsmontoya/blueprint"
	"github.com/missionMeteora/journaler"
)

func main() {
	var (
		//b   *blueprint.Blueprint
		err error
	)

	if _, err = blueprint.New("Test app", 640, 480, blueprint.NilColor); err != nil {
		journaler.Error("Error initializing blueprint: %v", err)
	}

	//b.Run()
}
