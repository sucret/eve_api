package event

type FooEvent struct {
}

func (f *FooEvent) Name() string {
	return "foo_event"
}
