package utils

type EventEmitter interface {
	Emit(key string, params ...interface{})
	On(key string, fn func(params ...interface{})) (id int32)
}

type eventEmitter struct {
	subscribers map[string][]func(params ...interface{})
}

func NewEventEmitter() EventEmitter {
	return &eventEmitter{
		subscribers: make(map[string][]func(params ...interface{})),
	}
}

func (e *eventEmitter) Emit(key string, params ...interface{}) {
	for _, subscriber := range e.subscribers[key] {
		subscriber(params...)
	}
}

func (e *eventEmitter) On(key string, fn func(params ...interface{})) int32 {
	e.subscribers[key] = append(e.subscribers[key], fn)

	return int32(len(e.subscribers[key]) - 1)
}
