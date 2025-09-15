package service

import (
	"errors"

	"github.com/ShivamSingh8853/go_event_driven/internal/events"
	"github.com/ShivamSingh8853/go_event_driven/internal/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var userStore = map[string]model.User{}

func RegisterUser(email, password string) error {
	_, exists := userStore[email]
	if exists {
		return errors.New("user already exists")
	}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := model.User{
		ID:       uuid.NewString(),
		Email:    email,
		Password: string(hashed),
	}
	userStore[email] = user

	events.Publish(events.UserCreatedEvents{Email: email})
	return nil
}

func LoginUser(email, password string) error {
	user, ok := userStore[email]
	if !ok {
		return errors.New("user not found")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New("invalid credentials")
	}
	return nil
}
