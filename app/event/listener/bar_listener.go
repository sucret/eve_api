package listener

import (
	appEvent "eve/app/event/event"
	"eve/internal/event"
	"fmt"
)

type BarListener struct{}

func (f BarListener) Listen() []event.EventInterface {
	return []event.EventInterface{
		&appEvent.FooEvent{},
	}
}

func (f BarListener) Process(e event.EventInterface) {
	fmt.Println("bar listener process event:", e, e.Name())
}
