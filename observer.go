package observer

type Event[T any] struct {
	Name string
}

func (e Event[T]) String() string {
	return e.Name
}

type EventProducer struct {
	observers map[string][]func(any)
}

func NewEventProducer() *EventProducer {
	return &EventProducer{observers: make(map[string][]func(any))}
}

func RegisterEvent[T any](p *EventProducer, evt Event[T], handler func(T)) {
	wrapped := func(d any) { handler(d.(T)) }

	p.observers[evt.Name] = append(p.observers[evt.Name], wrapped)
}

func FireEvent[T any](p *EventProducer, evt Event[T], data T) {
	for _, h := range p.observers[evt.Name] {
		h(data)
	}
}
