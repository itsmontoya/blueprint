package blueprint

import (
	"testing"
)

func TestBasic(t *testing.T) {
	var (
		b   *Blueprint
		err error
	)

	if b, err = New("Test app", 640, 480, NilColor); err != nil {
		t.Fatal(err)
	}

	b.Run()
}
