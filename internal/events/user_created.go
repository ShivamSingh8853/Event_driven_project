package events

type UserCreatedEvents struct {
	Email string
}

func (e UserCreatedEvents) Name() string {
	return "userCreated"
}
