package event

import (
	"eve/app/event/listener"
	"eve/internal/event"
	"eve/internal/global"
)

var listenerList = []event.ListenerInterface{
	listener.FooListener{},
	listener.BarListener{},
}

func init() {
	for _, l := range listenerList {
		global.EventDispatcher.Register(l)
	}
}
