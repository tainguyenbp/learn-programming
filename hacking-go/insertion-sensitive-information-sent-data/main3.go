package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/login", login)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	// Ensure we're dealing with a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		return
	}

	// Retrieve username and password from the form
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Authenticate the user
	if !authenticate(username, password) {
		log.Println("Login failed for user:", username)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Proceed with successful login
	// Here you might set a session or token
	log.Println("Login successful for user:", username)
	fmt.Fprintf(w, "Welcome, %s!", username)
}

func authenticate(username, password string) bool {
	// Replace with actual authentication logic
	// For example, compare with hardcoded credentials
	if username == "admin" && password == "password" {
		return true
	}
	return false
}
