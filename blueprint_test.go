package blueprint

import (
	"testing"
)

func TestBasic(t *testing.T) {
	var (
		b   *Blueprint
		err error
	)

	if b = New("Test app", Rects{Width: 640, Height: 480}, NilColor); err != nil {
		t.Fatal(err)
	}

	if err = b.Run(func() {}); err != nil {
		t.Fatal(err)
	}
}
