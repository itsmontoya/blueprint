package blueprint

import (
	"github.com/Path94/atoms"
	"github.com/missionMeteora/toolkit/errors"
)

// EventType represents an event type
type EventType uint8

const (
	// EventNil is the zero value for events
	EventNil EventType = iota
	// EventMouseEnter represents a mouse entering an element
	EventMouseEnter
	// EventMouseLeave represents a mouse leaving an element
	EventMouseLeave

	// EventMouseDown represents a mouse pressing down over an element
	EventMouseDown
	// EventMouseUp represents a mouse unpressing while over an element
	EventMouseUp
	// EventMouseClick represents a mouse pressing and unpressing over the same element
	EventMouseClick

	// EventOnFocus represents an element being focused
	EventOnFocus
	// EventOnBlur represents an element's focus being removed
	EventOnBlur
)

const (
	// ErrInvalidEvent is returned when an event is referenced which is not specified in the EventType const block
	ErrInvalidEvent = errors.Error("invalid event")
)

// NewEvents returns a new instance of Events
func NewEvents() *Events {
	var e Events
	return &e
}

// Events represents the supported system events
type Events struct {
	mux atoms.RWMux

	onMouseEnter *eventMap
	onMouseLeave *eventMap

	onMouseDown *eventMap
	onMouseUp   *eventMap

	onFocus *eventMap
	onBlur  *eventMap
}

func (e *Events) get(et EventType, create bool) (em *eventMap, err error) {
	switch et {
	case EventMouseEnter:
		if e.onMouseEnter == nil && create {
			e.onMouseEnter = newEventMap()
		}

		em = e.onMouseEnter

	case EventMouseLeave:
		if e.onMouseEnter == nil && create {
			e.onMouseEnter = newEventMap()
		}

		em = e.onMouseEnter

	case EventMouseDown:
		if e.onMouseEnter == nil && create {
			e.onMouseEnter = newEventMap()
		}

		em = e.onMouseEnter

	case EventMouseUp:
		if e.onMouseEnter == nil && create {
			e.onMouseEnter = newEventMap()
		}

		em = e.onMouseEnter

	case EventOnFocus:
		if e.onMouseEnter == nil && create {
			e.onMouseEnter = newEventMap()
		}

		em = e.onMouseEnter

	case EventOnBlur:
		if e.onMouseEnter == nil && create {
			e.onMouseEnter = newEventMap()
		}

		em = e.onMouseEnter

	default:
		err = ErrInvalidEvent
	}

	return
}

// Will return whether or not event type has any subscribers
func (e *Events) has(et EventType) (has bool) {
	e.mux.Read(func() {
		em, _ := e.get(et, false)
		has = em != nil
	})

	return
}

func (e *Events) notify(evt Event) (has bool) {
	var em *eventMap
	e.mux.Read(func() {
		if em, _ = e.get(evt.et, false); em == nil {
			return
		}

		// TODO: Once everything is stable, make this call notify within a goroutine
		has = em.notify(evt)
	})

	return
}

// Subscribe will subscribe to an event
func (e *Events) Subscribe(et EventType, fn EventFn) (key int, err error) {
	var em *eventMap
	e.mux.Update(func() {
		if em, err = e.get(et, true); err != nil {
			return
		}

		key = em.Subscribe(fn)
	})

	return
}

// Unsubscribe will unsubscribe from an event
func (e *Events) Unsubscribe(et EventType, key int) {
	var (
		em  *eventMap
		err error
	)

	e.mux.Update(func() {
		if em, err = e.get(et, false); em == nil || err != nil {
			return
		}

		em.Unsubscribe(key)
	})

	return
}

// Event represents an event
type Event struct {
	// Event type
	et EventType
	// Window position
	wp Coords
}

// EventFn is used as a callback for listening events
type EventFn func(Event)

func newEventMap() *eventMap {
	var em eventMap
	em.m = make(map[int]EventFn)
	return &em
}

type eventMap struct {
	mux atoms.RWMux
	// Key index
	ki int
	// Internal map
	m map[int]EventFn
}

func (em *eventMap) notify(evt Event) (has bool) {
	em.mux.Read(func() {
		if len(em.m) == 0 {
			return
		}

		for _, fn := range em.m {
			fn(evt)
		}

		has = true
	})

	return
}

func (em *eventMap) Subscribe(fn EventFn) (key int) {
	em.mux.Update(func() {
		key = em.ki
		em.ki++

		em.m[key] = fn
	})

	return
}

func (em *eventMap) Unsubscribe(key int) {
	em.mux.Update(func() {
		delete(em.m, key)
	})

	return
}
