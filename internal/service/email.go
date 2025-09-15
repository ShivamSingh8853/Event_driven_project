package service

import (
	"fmt"

	"github.com/ShivamSingh8853/go_event_driven/internal/events"
)

func InitEmailService() {
	events.Subscribe("userCreated", func(e events.Event) {
		event := e.(events.UserCreatedEvents)
		sendWelcomeEmail(event.Email)
	})
}
func sendWelcomeEmail(email string) {
	fmt.Printf("Sending welcome email to %s\n", email)
}
