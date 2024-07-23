package test

import (
	"eve/app/event/event"
	"eve/internal/global"
	"testing"
)

func TestEvent(t *testing.T) {
	global.EventDispatcher.Dispatch(&event.FooEvent{})
}
