package event

type ListenerInterface interface {
	Listen() []EventInterface
	Process(EventInterface)
}
