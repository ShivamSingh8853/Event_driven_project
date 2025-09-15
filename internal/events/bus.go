package events

type Event interface {
	Name() string
}

type Handler func(Event)

var subscribers = map[string][]Handler{}

func Subscribe(eventName string, handler Handler) {
	subscribers[eventName] = append(subscribers[eventName], handler)
}
func Publish(event Event) {
	for _, handler := range subscribers[event.Name()] {

		go handler(event)
	}
}
