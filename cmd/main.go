package main

import (
	"fmt"
	"net/http"

	"github.com/ShivamSingh8853/go_event_driven/internal/api"
	"github.com/ShivamSingh8853/go_event_driven/internal/service"
)

func main() {
	service.InitEmailService()
	http.HandleFunc("/register", api.RegisterHandler)
	http.HandleFunc("/login", api.LoginHandler)
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
