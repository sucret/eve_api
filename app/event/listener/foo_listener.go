package listener

import (
	appEvent "eve/app/event/event"
	"eve/internal/event"
	"fmt"
)

type FooListener struct{}

func (f FooListener) Listen() []event.EventInterface {
	return []event.EventInterface{
		&appEvent.FooEvent{},
	}
}

func (f FooListener) Process(e event.EventInterface) {
	fmt.Println("foo listener process event:", e, e.Name())
}
