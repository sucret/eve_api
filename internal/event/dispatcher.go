package event

import "sync"

type Dispatcher struct {
	listenerMap map[string][]ListenerInterface
	mu          sync.RWMutex
}

func New() *Dispatcher {
	return &Dispatcher{
		listenerMap: make(map[string][]ListenerInterface),
		mu:          sync.RWMutex{},
	}
}

// Register 注册监听者
func (d *Dispatcher) Register(listener ListenerInterface) {
	d.mu.Lock()
	defer d.mu.Unlock()

	// 获取监听者监听的事件列表
	eventList := listener.Listen()

	// 记录事件和监听者的关系
	for _, event := range eventList {
		d.listenerMap[event.Name()] = append(d.listenerMap[event.Name()], listener)
	}
}

// Dispatch 触发事件
func (d *Dispatcher) Dispatch(event EventInterface) {
	// 获取事件的监听者列表
	listenerList, exist := d.listenerMap[event.Name()]
	if !exist {
		return
	}

	// 执行listener的process方法
	for _, listener := range listenerList {
		listener.Process(event)
	}
}
