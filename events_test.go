package blueprint

import (
	"testing"

	"github.com/missionMeteora/journaler"
)

func TestEvents(t *testing.T) {
	var (
		evts *Events
		ch   = make(chan int, 2)
	//	err  error
	)

	evts = NewEvents()
	evts.Subscribe(EventMouseDown, func(evt Event) {
		ch <- 2
	})

	evts.Subscribe(EventMouseDown, func(evt Event) {
		ch <- 1
	})

	go evts.notify(Event{et: EventMouseDown})

	var val int
	val += <-ch
	val += <-ch

	if val != 3 {
		t.Fatalf("invalid value, expected %v and received %v", 3, val)
	}
}
